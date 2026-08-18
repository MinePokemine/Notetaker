package helpers

import "net/http"

func CookieOrHeader(key string, r *http.Request) (string, error) {
	out := r.Header.Get(key)

	if out != "" {
		return out, nil
	}

	cookie, err := r.Cookie(key)

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
