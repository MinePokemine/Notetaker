package handlers_account

import (
	"crypto/rand"
	"net/http"

	data "github.com/MinePokemine/notetaker/Data"
	helpers "github.com/MinePokemine/notetaker/Helpers"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	username := r.Form.Get("username")
	if username == "" {
		http.Error(w, "Null username", http.StatusBadRequest)
	}

	var login [8]byte
	rand.Reader.Read(login[:])

	cookie := &http.Cookie{
		Name:  "auth",
		Value: string(login[:]),
	}

	http.SetCookie(w, cookie)

	data.Users = append(data.Users, &helpers.User{
		Username: username,
		Login:    login,
		Projects: []*helpers.Project{},
	})
}
