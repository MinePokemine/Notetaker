package handlers_projects

import (
	"net/http"
	"strconv"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func Project(w http.ResponseWriter, r *http.Request) (t.ProjectEx, bool) {
	uid, _ := helpers_account.Auth(w, r)
	if uid < 0 {
		return t.ProjectEx{}, false
	}
	user := t.Users[uid]

	projIDStr := r.PathValue("pid")
	projID, err := strconv.Atoi(projIDStr)

	if err != nil {
		http.Error(w, "Project ID not an integer: "+err.Error(), http.StatusBadRequest)
		return t.ProjectEx{}, false
	}

	if projID > len(user.Projects) {
		http.Error(w, "Invalid project id", http.StatusBadRequest)
		return t.ProjectEx{}, false
	}

	return user.Projects[projID].Expand(), true
}
