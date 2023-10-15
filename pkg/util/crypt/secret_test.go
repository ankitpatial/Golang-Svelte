/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package crypt

import "testing"

func TestAccountKey(t *testing.T) {
	if _, err := AccountKey(); err != nil {
		t.Error(err)
		return
	}

	if _, err := AccountKey(); err != nil {
		t.Error(err)
		return
	}

	if _, err := AccountKey(); err != nil {
		t.Error(err)
		return
	}
}

func TestNewAesKey(t *testing.T) {
	hashKey := NewKey(64)
	if hashKey == "" {
		t.Error("hashKey is empty")
		return
	}

	t.Log("hashKey", string(hashKey))

	blockKey := NewKey(32)
	if blockKey == "" {
		t.Error("blockKey  is empty")
	}

	t.Log("blockKey", string(blockKey))
}
