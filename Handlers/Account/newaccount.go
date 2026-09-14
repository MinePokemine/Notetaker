package handlers_account

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"

	data "github.com/MinePokemine/notetaker/Data"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
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

	data.Users = append(data.Users, &t.User{
		Username: username,
		Login:    login,
		Projects: []*t.Project{},
		UID:      uid,
	})

	http.Redirect(w, r, "/account", http.StatusSeeOther)
}
