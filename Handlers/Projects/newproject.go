package handlers_projects

import (
	"net/http"
	"strconv"

	data "github.com/MinePokemine/notetaker/Data"
	helpers "github.com/MinePokemine/notetaker/Helpers"
	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
)

func NewProject(w http.ResponseWriter, r *http.Request) {
	uid, _ := helpers_account.Auth(w, r)
	if uid < 0 {
		return
	}
	user := data.Users[uid]

	projName := r.FormValue("name")
	if projName == "" {
		http.Error(w, "Project name not supplied", http.StatusBadRequest)
	}
	projID := len(user.Projects)

	user.Projects = append(user.Projects, &helpers.Project{
		Name:     projName,
		User:     user,
		IDInUser: projID,
	})

	http.Redirect(w, r, "/projects/"+strconv.Itoa(projID), http.StatusSeeOther)
}
