package helpers_account

import (
	"net/http"
	"strconv"

	data "github.com/MinePokemine/notetaker/Data"
	helpers "github.com/MinePokemine/notetaker/Helpers"
)

func Auth(w http.ResponseWriter, r *http.Request) (*helpers.User, error) {
	uidStr, err := helpers.CookieOrHeader("uid", w, r)
	if err != nil {
		http.Error(w, "UID not supplied", http.StatusBadRequest)
		return nil, err
	}

	authStr, err := helpers.CookieOrHeader("auth", w, r)
	if err != nil {
		http.Error(w, "UID not supplied", http.StatusBadRequest)
		return nil, err
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		http.Error(w, "UID not a number", http.StatusBadRequest)
		return nil, err
	}

	var auth [8]byte
	copy(auth[:], authStr)

	if !(len(data.Users) > uid) && data.Users[uid].Login == auth {
		http.Error(w, "Invalid user id or authentication", http.StatusUnauthorized)
		return nil, nil
	}

	return data.Users[uid], nil
}
