package handlers_notes

import (
	"net/http"
	"strconv"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func NewTag(w http.ResponseWriter, r *http.Request) {
	uid, err := helpers_account.Auth(w, r)
	if uid < 0 {
		return
	}

	user := t.Users[uid]

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

	var parents []t.TagReference

	tag := &t.Tag{
		UID: project.UID,
		PID: project.PID,
		TID: len(project.Tags),

		Name: name,
		//Tags:        tags,
		//Source:      source,
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

		parents = append(parents, t.TagReference{
			UID: project.UID,
			PID: project.PID,
			TID: tID,
		})

		project.Tags[tID].Children = append(project.Tags[tID].Children, tag.Reference())
	}

	tag.Parents = parents

	project.Tags = append(project.Tags, tag)

	http.Redirect(w, r, "/projects/"+pIDstr, http.StatusSeeOther)
}
