package handlers_account

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"

	data "github.com/MinePokemine/notetaker/Data"
	helpers "github.com/MinePokemine/notetaker/Helpers"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Null username", http.StatusBadRequest)
	}

	var login [32]byte
	rand.Reader.Read(login[:])

	uid := len(data.Users)

	cookie := &http.Cookie{
		Name:  "auth",
		Value: base64.RawURLEncoding.EncodeToString(login[:]),
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	cookie = &http.Cookie{
		Name:  "uid",
		Value: strconv.Itoa(uid),
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	data.Users = append(data.Users, &helpers.User{
		Username: username,
		Login:    login,
		Projects: []*helpers.Project{},
		UID:      uid,
	})

	http.Redirect(w, r, "/account", http.StatusSeeOther)
}
