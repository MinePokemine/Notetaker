package handlers_notes

import (
	"net/http"
	"strconv"

	data "github.com/MinePokemine/notetaker/Data"
	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func NewTag(w http.ResponseWriter, r *http.Request) {
	uid, err := helpers_account.Auth(w, r)
	if uid < 0 {
		return
	}

	user := data.Users[uid]

	pIDstr := r.PathValue("pid")
	pID, err := strconv.Atoi(pIDstr)
	if err != nil {
		http.Error(w, "Project ID not an int: "+err.Error(), http.StatusBadRequest)
		return
	}
	if pID > len(user.Projects) {
		http.Error(w, "Invalid project "+pIDstr, http.StatusBadRequest)
		return
	}
	project := user.Projects[pID]

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	parStrs := r.Form["parents"]
	name := r.FormValue("name")

	var parents []*t.Tag

	tag := &t.Tag{
		Name: name,
		//Tags:        tags,
		//Source:      source,
		IDInProject: len(project.Tags),
		Project:     project,
	}

	for i, tagStr := range parStrs {
		tID, err := strconv.Atoi(tagStr)
		if err != nil {
			http.Error(w, "Tag #"+strconv.Itoa(i)+" in the form not an int: "+err.Error(), http.StatusBadRequest)
			return
		}
		if tID > len(project.Tags) {
			println("Invalid tag (Tag #" + strconv.Itoa(i) + " in form): " + strconv.Itoa(tID))
			continue
		}

		parents = append(parents, project.Tags[tID])

		project.Tags[tID].Parents = append(project.Tags[tID].Parents, tag)
	}

	tag.Parents = parents

	project.Tags = append(project.Tags, tag)

	http.Redirect(w, r, "/projects/"+pIDstr, http.StatusSeeOther)
}
