package handlers_account

import (
	"net/http"

	data "github.com/MinePokemine/notetaker/Data"
	helpers "github.com/MinePokemine/notetaker/Helpers"
	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
)

func MyAccount(w http.ResponseWriter, r *http.Request) (helpers.User, bool) {
	uid, _ := helpers_account.Auth(w, r)

	if uid == -1 {
		return helpers.User{}, false
	}

	return *data.Users[uid], true
}
