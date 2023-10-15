package oauth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// GetBasicAuthentication get username and password from Authorization header
func GetBasicAuthentication(r *http.Request) (username, password string, err error) {
	if header := r.Header.Get("Authorization"); header != "" {
		if strings.ToLower(header[:6]) == "basic " {
			// decode header value
			value, err := base64.StdEncoding.DecodeString(header[6:])
			if err != nil {
				return "", "", err
			}
			strValue := string(value)
			if ind := strings.Index(strValue, ":"); ind > 0 {
				return strValue[:ind], strValue[ind+1:], nil
			}
		}
	}
	return "", "", nil
}
