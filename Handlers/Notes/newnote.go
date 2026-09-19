package handlers_notes

import (
	"net/http"
	"strconv"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func NewNote(w http.ResponseWriter, r *http.Request) {
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

	tagStrs := r.Form["tags"]
	data := r.FormValue("data")

	var tags []t.TagReference

	note := &t.Note{
		UID: project.UID,
		PID: project.PID,
		NID: len(project.Notes),

		Data: data,
		//Tags:        tags,
		//Source:      source,
	}

	for i, tagStr := range tagStrs {
		tID, err := strconv.Atoi(tagStr)
		if err != nil {
			http.Error(w, "Tag #"+strconv.Itoa(i)+" in the form not an int: "+err.Error(), http.StatusBadRequest)
			return
		}
		if tID > len(project.Tags) {
			println("Invalid tag (Tag #" + strconv.Itoa(i) + " in form): " + strconv.Itoa(tID))
			continue
		}

		tags = append(tags, project.Tags[tID].Reference())

		project.Tags[tID].Notes = append(project.Tags[tID].Notes, note.Reference())
	}

	source := r.FormValue("src")
	note.Source = source
	note.Tags = tags

	project.Notes = append(project.Notes, note)

	http.Redirect(w, r, "/projects/"+pIDstr, http.StatusSeeOther)
}
