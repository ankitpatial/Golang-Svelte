/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package crypt

import (
	"context"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"roofix/config"
	"roofix/pkg/util/storage"
)

type PemKey string
type AesKey string

var cache = make(map[string]interface{})

const (
	folder               = "secret"
	accountSecret PemKey = "account-secret"
	appSecret     AesKey = "app-secret"
)

func AccountKey() (*rsa.PrivateKey, error) {
	return rsaKey(context.Background(), accountSecret)
}

func readSecret(keyName AesKey) ([]byte, error) {
	name := fmt.Sprintf("%s.txt", keyName)
	if k, ok := cache[string(keyName)]; ok {
		return k.([]byte), nil
	}

	// read from server
	bucket := config.DataBucket()
	object := fmt.Sprintf("%s/%s", folder, name)
	if data, e := storage.GetObject(context.Background(), bucket, object); e != nil {
		return nil, e
	} else if data == nil {
		// generate key
		key, err := NewAesKey(32)
		if err != nil {
			return nil, err
		}

		// save to s3
		ct := "text/plain"
		hexStr := hex.EncodeToString(key)
		err = storage.PutObject(context.Background(), bucket, object, []byte(hexStr), &ct)
		if err != nil {
			return nil, err
		}

		// set to cache & return
		cache[string(keyName)] = key
		return key, nil
	} else {
		d, _ := hex.DecodeString(string(data))

		// set to cache & return
		cache[string(keyName)] = d
		return d, nil
	}
}
