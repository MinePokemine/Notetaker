package handlers_projects

import (
	"net/http"
	"strconv"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func NewProject(w http.ResponseWriter, r *http.Request) {
	uid, _ := helpers_account.Auth(w, r)
	if uid < 0 {
		return
	}
	user := t.Users[uid]

	projName := r.FormValue("name")
	if projName == "" {
		http.Error(w, "Project name not supplied", http.StatusBadRequest)
	}
	projID := len(user.Projects)

	user.Projects = append(user.Projects, &t.Project{
		UID: user.UID,
		PID: projID,

		Name: projName,
	})

	http.Redirect(w, r, "/projects/"+strconv.Itoa(projID), http.StatusSeeOther)
}
