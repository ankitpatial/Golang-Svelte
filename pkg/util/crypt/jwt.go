/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package crypt

import (
	"crypto/rsa"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"roofix/config"
	"time"
)

func JWTSign(data map[string]interface{}, d time.Duration) (string, error) {
	// gen validate token
	key, err := AccountKey()
	if err != nil {
		return "", err
	}

	return jwtToken(key, data, d)
}

func JWTParse(token string) (map[string]interface{}, error) {
	// gen validate token
	key, err := AccountKey()
	if err != nil {
		return nil, err
	}

	t, err := jwtParseToken(key, token)
	if err != nil {
		return nil, err
	}

	return t.PrivateClaims(), nil
}

func jwtToken(rsaKey *rsa.PrivateKey, claims map[string]interface{}, d time.Duration) (string, error) {
	prv, err := jwk.FromRaw(rsaKey)
	if err != nil {
		return "", fmt.Errorf("failed to create JWK, %w", err)
	}

	builder := jwt.NewBuilder().
		Issuer(config.Read().Website.Domain()).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(d))

	for k, v := range claims {
		builder = builder.Claim(k, v)
	}

	token, err := builder.Build()
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, prv))
	if err != nil {
		return "", fmt.Errorf("failed to generate signed payload, %w", err)
	}

	return string(signed), nil
}

func jwtParseToken(rsaKey *rsa.PrivateKey, payload string) (jwt.Token, error) {
	prv, err := jwk.FromRaw(rsaKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWK, %w", err)
	}

	pub, err := jwk.PublicKeyOf(prv)
	if err != nil {
		return nil, fmt.Errorf("failed on jwk.FromRaw, %w", err)
	}

	token, err := jwt.Parse([]byte(payload), jwt.WithKey(jwa.RS256, pub))
	if err != nil {
		return nil, fmt.Errorf("failed on jwk.FromRaw, %w", err)
	}

	return token, nil
}
