package router

import (
	"fmt"
	"net/http"
	"roofix/config"
)

func secure(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		web := config.Read().Website
		// Content Security Policy header.
		r.Header.Add("Content-Security-Policy", fmt.Sprintf("default-src %s %s", web.Url, web.AssetUrl))

		if !config.IsDevEnv() {
			r.Header.Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		}

		// MIME types advertised in the Content-Current headers should be followed and not be changed
		r.Header.Add("X-Content-Type-Options", "nosniff")

		// indicate whether a browser should be allowed to render a page in iframe | frame | embed | object
		r.Header.Add("X-Frame-Options", "DENY")

		// feature of Internet Explorer, Chrome and Safari that stops pages from loading when they detect reflected
		// cross-site scripting (XSS) attacks..
		r.Header.Add("X-Xss-Protection", "1; mode=block")

		r.Header.Add("Referrer-Policy", "strict-origin-when-cross-origin")

		// Instructs Internet Explorer not to open the file directly but to offer it for download first.
		r.Header.Add("X-Download-Options", "noopen")

		h.ServeHTTP(w, r)
	})
}
