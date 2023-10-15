package oauth

import (
	"context"
	"errors"
	"net/http"
	"roofix/ent"
	"roofix/pkg/external"
)

// Verify provides user credentials verifier for testing.
type Verify struct {
}

// ValidateUser validates username and password returning an error if the user credentials are wrong
func (*Verify) ValidateUser(username, password, scope string, r *http.Request) (string, error) {
	return external.AuthenticateApiUser(r.Context(), username, password)
}

// ValidateClient validates clientID and secret returning an error if the client credentials are wrong
func (*Verify) ValidateClient(clientID, clientSecret, scope string, r *http.Request) (string, error) {
	return external.AuthenticateApiUser(r.Context(), clientID, clientSecret)
}

// ValidateCode validates token ID
func (*Verify) ValidateCode(clientID, clientSecret, code, redirectURI string, r *http.Request) (string, error) {
	return "", nil
}

// AddClaims provides additional claims to the token
func (*Verify) AddClaims(tokenType TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	claims := make(map[string]string)
	u := external.GetApiUser(r.Context(), credential)
	if u != nil {
		claims["apiUID"] = u.ID
		claims["apiUName"] = u.Username
	}
	return claims, nil
}

// StoreTokenID saves the token id generated for the user
func (*Verify) StoreTokenID(tokenType TokenType, credential, tokenID, refreshTokenID string) error {
	return external.SaveApiUserToken(context.Background(), credential, string(tokenType), tokenID, refreshTokenID)
}

// AddProperties provides additional information to the token response
func (*Verify) AddProperties(tokenType TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	return nil, nil
}

// ValidateTokenID validates token ID
func (*Verify) ValidateTokenID(tokenType TokenType, credential, tokenID, refreshTokenID string) (string, error) {
	ctx := context.Background()
	// check if user is active
	u := external.GetApiUser(ctx, credential)
	if u == nil || !u.Active {
		return "", errors.New("user not exist or in-active")
	}

	// check if token exists
	if _, err := external.GetApiUserToken(ctx, tokenID); err != nil {
		if ent.IsNotFound(err) {
			return "", errors.New("token not exist")
		}

		return "", err
	}

	return u.ID, nil
}
