package handlers_account

import (
	"net/http"

	data "github.com/MinePokemine/notetaker/Data"
	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
)

func RenameAccount(w http.ResponseWriter, r *http.Request) {
	uid, _ := helpers_account.Auth(w, r)
	if uid == -1 {
		return
	}

	data.Users[uid].Username = r.FormValue("username")

	http.Redirect(w, r, "/account/", http.StatusSeeOther)
}
