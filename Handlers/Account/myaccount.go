package handlers_account

import (
	"net/http"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func MyAccount(w http.ResponseWriter, r *http.Request) (t.User, bool) {
	uid, _ := helpers_account.Auth(w, r)

	if uid == -1 {
		return t.User{}, false
	}

	return *t.Users[uid], true
}
