package helpers_account

import (
	"encoding/base64"
	"net/http"
	"strconv"

	data "github.com/MinePokemine/notetaker/Data"
	helpers "github.com/MinePokemine/notetaker/Helpers"
)

func Auth(w http.ResponseWriter, r *http.Request) (int, error) {
	uidStr, err := helpers.CookieOrHeader("uid", r)
	if err != nil {
		http.Error(w, "UID not supplied", http.StatusBadRequest)
		return -1, err
	}

	authStr, err := helpers.CookieOrHeader("auth", r)
	if err != nil {
		http.Error(w, "Authentication not supplied", http.StatusBadRequest)
		return -1, err
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		http.Error(w, "UID not a number", http.StatusBadRequest)
		return -1, err
	}

	authSlice, err := base64.RawURLEncoding.DecodeString(authStr)
	if err != nil {
		http.Error(w, "Authentication string not a base64 Raw URL Encoding", http.StatusBadRequest)
	}
	auth := [32]byte(authSlice)

	if len(data.Users) < uid && data.Users[uid].Login == auth {
		http.Error(w, "Invalid user id or authentication", http.StatusUnauthorized)
		return -1, nil
	}

	return uid, nil
}
