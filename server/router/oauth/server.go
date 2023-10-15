package oauth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"roofix/pkg/audit"
	"roofix/pkg/util/uid"
	"roofix/server/router/response"
	"time"
)

type GrantType string

const (
	PasswordGrant          GrantType = "password"
	ClientCredentialsGrant GrantType = "client_credentials"
	AuthCodeGrant          GrantType = "authorization_code"
	RefreshTokenGrant      GrantType = "refresh_token"
)

// CredentialsVerifier defines the interface of the user and client credentials verifier.
type CredentialsVerifier interface {
	// ValidateUser Validate username and password returning an error if the user credentials are wrong
	ValidateUser(username, password, scope string, r *http.Request) (string, error)
	// ValidateClient Validate clientID and secret returning an error if the client credentials are wrong
	ValidateClient(clientID, clientSecret, scope string, r *http.Request) (string, error)
	// AddClaims Provide additional claims to the token
	AddClaims(tokenType TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error)
	// AddProperties Provide additional information to the authorization server response
	AddProperties(tokenType TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error)
	// ValidateTokenID Optionally validate previously stored tokenID during refresh request
	ValidateTokenID(tokenType TokenType, credential, tokenID, refreshTokenID string) (string, error)
	// StoreTokenID Optionally store the tokenID generated for the user
	StoreTokenID(tokenType TokenType, credential, tokenID, refreshTokenID string) error
}

// AuthorizationCodeVerifier defines the interface of the Authorization Code verifier
type AuthorizationCodeVerifier interface {
	// ValidateCode checks the authorization code and returns the user credential
	ValidateCode(clientID, clientSecret, code, redirectURI string, r *http.Request) (string, error)
}

// BearerServer is the OAuth 2 bearer server implementation.
type BearerServer struct {
	secretKey string
	TokenTTL  time.Duration
	verifier  CredentialsVerifier
	provider  *TokenProvider
}

// NewBearerServer creates new OAuth 2 bearer server
func NewBearerServer(secretKey string, ttl time.Duration, verifier CredentialsVerifier, formatter TokenSecureFormatter) *BearerServer {
	if formatter == nil {
		formatter = NewSHA256RC4TokenSecurityProvider([]byte(secretKey))
	}
	return &BearerServer{
		secretKey: secretKey,
		TokenTTL:  ttl,
		verifier:  verifier,
		provider:  NewTokenProvider(formatter)}
}

func (bs *BearerServer) ApiUserCred(w http.ResponseWriter, r *http.Request) {
	var grantType, username, password, refreshToken string
	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		grantType = r.FormValue("grant_type")
		username = r.FormValue("username")
		password = r.FormValue("password")
		refreshToken = r.FormValue("refresh_token")
		if len(refreshToken) > 0 { // decode token
			b, err := base64.RawURLEncoding.DecodeString(r.FormValue("refresh_token"))
			if err != nil {
				refreshToken = string(b)
			}
		}
	default:
		var inp map[string]string
		if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
			response.BadRequest(w, err)
			return
		}
		grantType, _ = inp["grant_type"]
		username, _ = inp["username"]
		password, _ = inp["password"]
		refreshToken, _ = inp["refresh_token"]
	}

	scope := "ApiClient"

	if username == "" || password == "" {
		// get username and password from basic authorization header
		var err error
		username, password, err = GetBasicAuthentication(r)
		if err != nil {
			renderJSON(w, "Not authorized", http.StatusUnauthorized)
			return
		}
	}

	resp, statusCode := bs.generateTokenResponse(GrantType(grantType), username, password, refreshToken, scope, "", "", r)
	renderJSON(w, resp, statusCode)
}

// UserCredentials manages password grant type requests
func (bs *BearerServer) UserCredentials(w http.ResponseWriter, r *http.Request) {
	grantType := r.FormValue("grant_type")

	username := r.FormValue("username")
	password := r.FormValue("password")
	scope := r.FormValue("scope")

	if username == "" || password == "" {
		// get username and password from basic authorization header
		var err error
		username, password, err = GetBasicAuthentication(r)
		if err != nil {
			renderJSON(w, "Not authorized", http.StatusUnauthorized)
			return
		}
	}

	refreshToken := r.FormValue("refresh_token")
	resp, statusCode := bs.generateTokenResponse(GrantType(grantType), username, password, refreshToken, scope, "", "", r)
	renderJSON(w, resp, statusCode)
}

// ClientCredentials manages client credentials grant type requests
func (bs *BearerServer) ClientCredentials(w http.ResponseWriter, r *http.Request) {
	grantType := r.FormValue("grant_type")
	// grant_type client_credentials variables
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	if clientID == "" || clientSecret == "" {
		// get clientID and secret from basic authorization header
		var err error
		clientID, clientSecret, err = GetBasicAuthentication(r)
		if err != nil {
			renderJSON(w, "Not authorized", http.StatusUnauthorized)
			return
		}
	}
	scope := r.FormValue("scope")
	refreshToken := r.FormValue("refresh_token")
	resp, statusCode := bs.generateTokenResponse(GrantType(grantType), clientID, clientSecret, refreshToken, scope, "", "", r)
	renderJSON(w, resp, statusCode)
}

// AuthorizationCode manages authorization code grant type requests for the phase two of the authorization process
func (bs *BearerServer) AuthorizationCode(w http.ResponseWriter, r *http.Request) {
	grantType := r.FormValue("grant_type")
	// grant_type client_credentials variables
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret") // not mandatory
	code := r.FormValue("code")
	redirectURI := r.FormValue("redirect_uri") // not mandatory
	scope := r.FormValue("scope")              // not mandatory
	if clientID == "" {
		var err error
		clientID, clientSecret, err = GetBasicAuthentication(r)
		if err != nil {
			renderJSON(w, "Not authorized", http.StatusUnauthorized)
			return
		}
	}
	resp, status := bs.generateTokenResponse(GrantType(grantType), clientID, clientSecret, "", scope, code, redirectURI, r)
	renderJSON(w, resp, status)
}

// Generate token response
func (bs *BearerServer) generateTokenResponse(grantType GrantType, credential string, secret string, refreshToken string, scope string, code string, redirectURI string, r *http.Request) (interface{}, int) {
	var resp *TokenResponse
	ctx := context.Background()
	switch grantType {
	case PasswordGrant:
		var apiUID string
		if id, err := bs.verifier.ValidateUser(credential, secret, scope, r); err != nil {
			audit.OpFailed(ctx, audit.OAuthPasswordGrant, err)
			return "Not authorized", http.StatusUnauthorized
		} else {
			apiUID = id
		}

		token, refresh, err := bs.generateTokens(UserToken, credential, scope, r)
		if err != nil {
			msg := "Token generation failed, check claims"
			audit.ApiOpFailed(ctx, audit.OAuthPasswordGrant, apiUID, errors.New(msg))
			return msg, http.StatusInternalServerError
		}

		if err = bs.verifier.StoreTokenID(token.TokenType, credential, token.ID, refresh.RefreshTokenID); err != nil {
			msg := "Storing Token ID failed"
			audit.ApiOpFailed(ctx, audit.OAuthPasswordGrant, apiUID, errors.New(msg))
			return msg, http.StatusInternalServerError
		}

		if resp, err = bs.cryptTokens(token, refresh, r); err != nil {
			msg := "Token generation failed, check security provider"
			audit.ApiOpFailed(ctx, audit.OAuthPasswordGrant, apiUID, errors.New(msg))
			return msg, http.StatusInternalServerError
		}

		audit.ApiOpSuccess(ctx, audit.OAuthPasswordGrant, apiUID, fmt.Sprintf("granted token: %s", token.ID))

	case ClientCredentialsGrant:
		if _, err := bs.verifier.ValidateClient(credential, secret, scope, r); err != nil {
			return "Not authorized", http.StatusUnauthorized
		}

		token, refresh, err := bs.generateTokens(ClientToken, credential, scope, r)
		if err != nil {
			return "Token generation failed, check claims", http.StatusInternalServerError
		}

		if err = bs.verifier.StoreTokenID(token.TokenType, credential, token.ID, refresh.RefreshTokenID); err != nil {
			return "Storing Token ID failed", http.StatusInternalServerError
		}

		if resp, err = bs.cryptTokens(token, refresh, r); err != nil {
			return "Token generation failed, check security provider", http.StatusInternalServerError
		}
	case AuthCodeGrant:
		codeVerifier, ok := bs.verifier.(AuthorizationCodeVerifier)
		if !ok {
			return "Not authorized, grant type not supported", http.StatusUnauthorized
		}

		user, err := codeVerifier.ValidateCode(credential, secret, code, redirectURI, r)
		if err != nil {
			return "Not authorized", http.StatusUnauthorized
		}

		token, refresh, err := bs.generateTokens(AuthToken, user, scope, r)
		if err != nil {
			return "Token generation failed, check claims", http.StatusInternalServerError
		}

		err = bs.verifier.StoreTokenID(token.TokenType, user, token.ID, refresh.RefreshTokenID)
		if err != nil {
			return "Storing Token ID failed", http.StatusInternalServerError
		}

		if resp, err = bs.cryptTokens(token, refresh, r); err != nil {
			return "Token generation failed, check security provider", http.StatusInternalServerError
		}
	case RefreshTokenGrant:
		refresh, err := bs.provider.DecryptRefreshTokens(refreshToken)
		if err != nil {
			msg := "Not authorized"
			audit.OpFailed(ctx, audit.OAuthRefreshTokenGrant, errors.New(msg))
			return msg, http.StatusUnauthorized
		}

		var apiUID string
		if id, err := bs.verifier.ValidateTokenID(refresh.TokenType, refresh.Credential, refresh.TokenID, refresh.RefreshTokenID); err != nil {
			msg := "Not authorized invalid token"
			audit.OpFailed(ctx, audit.OAuthRefreshTokenGrant, errors.New(msg))
			return msg, http.StatusUnauthorized
		} else {
			apiUID = id
		}

		token, refresh, err := bs.generateTokens(refresh.TokenType, refresh.Credential, refresh.Scope, r)
		if err != nil {
			msg := "Token generation failed"
			audit.ApiOpFailed(ctx, audit.OAuthRefreshTokenGrant, apiUID, errors.New(msg))
			return msg, http.StatusInternalServerError
		}

		err = bs.verifier.StoreTokenID(token.TokenType, refresh.Credential, token.ID, refresh.RefreshTokenID)
		if err != nil {
			msg := "Storing Token ID failed"
			audit.ApiOpFailed(ctx, audit.OAuthRefreshTokenGrant, apiUID, errors.New(msg))
			return msg, http.StatusInternalServerError
		}

		if resp, err = bs.cryptTokens(token, refresh, r); err != nil {
			msg := "Token generation failed"
			audit.ApiOpFailed(ctx, audit.OAuthRefreshTokenGrant, apiUID, errors.New(msg))
			return msg, http.StatusInternalServerError
		}

		audit.ApiOpSuccess(ctx, audit.OAuthRefreshTokenGrant, apiUID, fmt.Sprintf("refresh token: %s", token.ID))

	default:
		return "Invalid grant_type", http.StatusBadRequest
	}

	return resp, http.StatusOK
}

func (bs *BearerServer) generateTokens(tokenType TokenType, username, scope string, r *http.Request) (*Token, *RefreshToken, error) {
	token := &Token{ID: uid.ULID(), Credential: username, ExpiresIn: bs.TokenTTL, CreationDate: time.Now().UTC(), TokenType: tokenType, Scope: scope}
	if bs.verifier != nil {
		claims, err := bs.verifier.AddClaims(token.TokenType, username, token.ID, token.Scope, r)
		if err != nil {
			return nil, nil, err
		}
		token.Claims = claims
	}

	refreshToken := &RefreshToken{RefreshTokenID: uid.ULID(), TokenID: token.ID, CreationDate: time.Now().UTC(), Credential: username, TokenType: tokenType, Scope: scope}

	return token, refreshToken, nil
}

func (bs *BearerServer) cryptTokens(token *Token, refresh *RefreshToken, r *http.Request) (*TokenResponse, error) {
	cToken, err := bs.provider.CryptToken(token)
	if err != nil {
		return nil, err
	}
	cRefreshToken, err := bs.provider.CryptRefreshToken(refresh)
	if err != nil {
		return nil, err
	}

	tokenResponse := &TokenResponse{Token: cToken, RefreshToken: cRefreshToken, TokenType: BearerToken, ExpiresIn: (int64)(bs.TokenTTL / time.Second)}

	if bs.verifier != nil {
		props, err := bs.verifier.AddProperties(token.TokenType, token.Credential, token.ID, token.Scope, r)
		if err != nil {
			return nil, err
		}
		tokenResponse.Properties = props
	}
	return tokenResponse, nil
}
