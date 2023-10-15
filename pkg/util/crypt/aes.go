/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Encrypt a string using AES
func Encrypt(str string) (string, error) {
	key, err := readSecret(appSecret)
	if err != nil {
		return "", err
	}

	gcm, err := getGCM(key)
	if err != nil {
		return "", err
	}

	// create a nonce. Nonce should be from GCM
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// encrypt the data using gcm.Seal
	// prefix nonce to the encrypted data.
	ciphertext := gcm.Seal(nonce, nonce, []byte(str), nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(encStr string) (string, error) {
	key, err := readSecret(appSecret)
	if err != nil {
		return "", err
	}

	// create a new Cipher Block from the key
	gcm, err := getGCM(key)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	enc, _ := hex.DecodeString(encStr)
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	// decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func getGCM(key []byte) (gcm cipher.AEAD, err error) {
	// cipher block from the key
	b, err := aes.NewCipher(key)
	if err != nil {
		return gcm, err
	}

	// GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	// https://golang.org/pkg/crypto/cipher/#NewGCM
	gcm, err = cipher.NewGCM(b)
	if err != nil {
		return gcm, err
	}

	return gcm, nil
}

func NewAesKey(length int) ([]byte, error) {
	bytes := make([]byte, length) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}

	return bytes, nil //encode key in bytes to string for saving
}

func NewKey(length int) string {
	k, err := NewAesKey(length)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(k)
}
