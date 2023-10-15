/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package crypt

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"roofix/config"
	"roofix/pkg/util/log"
	"roofix/pkg/util/storage"
)

func NewRsaKey() (*rsa.PrivateKey, error) {
	bits := 2048
	return rsa.GenerateKey(rand.Reader, bits)
}

func exportRsaKey(key *rsa.PrivateKey) []byte {
	raw := x509.MarshalPKCS1PrivateKey(key)
	pemKey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: raw,
		},
	)
	return pemKey
}

// ParsePrivateKey
func parseRsaKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func rsaKey(ctx context.Context, name PemKey) (*rsa.PrivateKey, error) {
	if k, ok := cache[string(name)]; ok {
		log.Info("using %s from cache", name)
		return k.(*rsa.PrivateKey), nil
	}

	bucket := config.DataBucket()
	key := fmt.Sprintf("%s/%s.pem", folder, name)
	data, e := storage.GetObject(ctx, bucket, key)
	if e != nil {
		return nil, e
	} else if data == nil {
		// generate secret
		pvk, err := NewRsaKey()
		if err != nil {
			return nil, err
		}

		// save to s3
		ct := "text/plain"
		err = storage.PutObject(ctx, bucket, key, exportRsaKey(pvk), &ct)
		if err != nil {
			return nil, err
		}

		// set to cache & return
		cache[string(name)] = pvk
		return pvk, nil
	}

	pvk, err := parseRsaKey(data)
	if err != nil {
		return nil, err
	}

	// set to cache & return
	cache[string(name)] = pvk
	return pvk, nil
}

//func ExportPublicKey(pubkey *rsa.PublicKey) (string, error) {
//	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
//	if err != nil {
//		return "", err
//	}
//	pubkey_pem := pem.EncodeToMemory(
//		&pem.Block{
//			Current:  "RSA PUBLIC KEY",
//			Bytes: pubkey_bytes,
//		},
//	)
//
//	return string(pubkey_pem), nil
//}
//
//func ParsePublicKey(pubPEM string) (*rsa.PublicKey, error) {
//	block, _ := pem.Decode([]byte(pubPEM))
//	if block == nil {
//		return nil, errors.New("failed to parse PEM block containing the key")
//	}
//
//	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
//	if err != nil {
//		return nil, err
//	}
//
//	switch pub := pub.(type) {
//	case *rsa.PublicKey:
//		return pub, nil
//	default:
//		break // fall through
//	}
//	return nil, errors.New("Key type is not RSA")
//}
