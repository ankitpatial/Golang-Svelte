package account

import (
	"roofix/pkg/util/crypt"
	"testing"
	"time"
)

func TestJwtToken(t *testing.T) {
	data := map[string]interface{}{
		"id":   123,
		"name": "ankit",
	}

	key, _ := crypt.NewRsaKey()

	// create token
	_, err := jwtToken(key, data, time.Second)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestJwtVerify(t *testing.T) {
	data := map[string]interface{}{
		"id":   123,
		"name": "ankit",
	}

	key, _ := crypt.NewRsaKey()

	// create token
	token, err := jwtToken(key, data, time.Second)
	if err != nil {
		t.Error(err)
		return
	}

	// verify token

	t2, err := jwtParseToken(key, token)
	if err != nil {
		t.Error(err)
		return
	}

	claims := t2.PrivateClaims()
	if _, ok := claims["id"]; !ok {
		t.Error("must have private claim['id']")
	}

	time.Sleep(time.Second)
	// must be expired by now
	_, err = jwtParseToken(key, token)
	if err == nil {
		t.Error("must have raised issue of expired token")
		return
	}

}
