package handlers_account

import (
	"net/http"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func RenameAccount(w http.ResponseWriter, r *http.Request) {
	uid, _ := helpers_account.Auth(w, r)
	if uid == -1 {
		return
	}

	t.Users[uid].Username = r.FormValue("username")

	http.Redirect(w, r, "/account/", http.StatusSeeOther)
}
