/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package crypt

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRandomKey(t *testing.T) {
	for i := 0; i < 1000; i++ {
		key, err := NewAesKey(32)
		if err != nil {
			t.Error(err)
			return
		}

		t.Log(fmt.Sprintf("%d: %v", i, hex.EncodeToString(key)))
	}
}

func TestEncryptDecrypt(t *testing.T) {
	s := "some string to test"
	Encrypt(s)
	Encrypt(s)
	enc, err := Encrypt(s)
	if err != nil {
		t.Error(err)
		return
	}

	d, err := Decrypt(enc)
	if err != nil {
		t.Error(err)
		return
	}

	if s != d {
		t.Error("expected: ", s, "but got:", d)
	}
}
