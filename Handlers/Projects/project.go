package handlers_projects

import (
	"net/http"
	"strconv"

	data "github.com/MinePokemine/notetaker/Data"
	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func Project(w http.ResponseWriter, r *http.Request) (t.Project, bool) {
	uid, _ := helpers_account.Auth(w, r)
	if uid < 0 {
		return t.Project{}, false
	}
	user := data.Users[uid]

	projIDStr := r.PathValue("pid")
	projID, err := strconv.Atoi(projIDStr)

	if err != nil {
		http.Error(w, "Project ID not an integer: "+err.Error(), http.StatusBadRequest)
		return t.Project{}, false
	}

	if projID > len(user.Projects) {
		http.Error(w, "Invalid project id", http.StatusBadRequest)
		return t.Project{}, false
	}

	return *user.Projects[projID], true
}
