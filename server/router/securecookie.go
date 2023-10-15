package router

import (
	"encoding/hex"
	"errors"
	"net/http"
	"roofix/config"
	"roofix/ent"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"time"
)

var cookies map[string]*crypt.SecureCookie

func init() {
	c := config.Read().Cookie
	h1, _ := hex.DecodeString(c.Previous.HashKey)
	b1, _ := hex.DecodeString(c.Previous.BlockKey)
	h2, _ := hex.DecodeString(c.Current.HashKey)
	b2, _ := hex.DecodeString(c.Current.BlockKey)
	cookies = map[string]*crypt.SecureCookie{
		"previous": crypt.NewSecureCookie(h1, b1),
		"current":  crypt.NewSecureCookie(h2, b2),
	}
}

func setSessionCookie(data string, expiresAt time.Time) *http.Cookie {
	name := conf.SessionCookieId
	c := &http.Cookie{
		Name:     name,
		Path:     "/",
		Domain:   conf.Website.Domain(),
		Expires:  expiresAt,
		HttpOnly: true,             // true: inaccessible to the JavaScript Document.cookie Website
		Secure:   !conf.IsDevEnv(), // true: request over the HTTPS protocol only
		SameSite: http.SameSiteStrictMode,
	}

	if encoded, err := crypt.EncodeMulti(name, data, cookies["current"]); err == nil {
		c.Value = encoded
	} else {
		c.Value = data
	}

	return c
}

func readSessionCookie(r *http.Request) string {
	name := conf.SessionCookieId
	cookie, err := r.Cookie(name)
	if err != nil {
		if !errors.Is(err, http.ErrNoCookie) {
			log.Error(err)
		}
		return ""
	}

	var value string
	err = crypt.DecodeMulti(name, cookie.Value, &value, cookies["current"], cookies["previous"])
	if err != nil {
		log.Error(err)
		return ""
	}

	return value
}

func clearSessionCookie(w http.ResponseWriter, r *http.Request) {
	// session cookie
	ctx := r.Context()
	sc, _ := r.Cookie(conf.SessionCookieId)
	if sc != nil && sc.Value != "" {
		var sid string
		err := crypt.DecodeMulti(conf.SessionCookieId, sc.Value, &sid, cookies["current"], cookies["previous"])
		if err != nil { // decrypt issue
			log.Error(err)
		} else { // remove from db
			if err := ent.GetClient().UserSession.DeleteOneID(sid).Exec(ctx); err != nil {
				log.Error(err)
			}
		}
	}

	// clear session cookie
	http.SetCookie(w, setSessionCookie("", time.Now()))
}
