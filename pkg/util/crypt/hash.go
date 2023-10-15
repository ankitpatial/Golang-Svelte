/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"math/big"
)

func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func MD5Int(text string) uint64 {
	n := new(big.Int)
	n.SetString(MD5Hash(text), 16)
	return n.Uint64()
}

func Hash(txt string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(txt), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hash)
}

func CompareHash(hash string, plainTxt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainTxt))
	return err == nil
}
