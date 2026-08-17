package handlers_account

import (
	"net/http"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
)

func RenameAccount(w http.ResponseWriter, r *http.Request) {
	usr, err := helpers_account.Auth(w, r)
	if err != nil || usr == nil {
		return
	}

	usr.Username = r.Form.Get("new")
}
