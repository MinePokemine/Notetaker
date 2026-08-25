package handlers_notes

import (
	"net/http"
	"strconv"

	data "github.com/MinePokemine/notetaker/Data"
	helpers "github.com/MinePokemine/notetaker/Helpers"
	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
)

func Note(w http.ResponseWriter, r *http.Request) (*helpers.Note, bool) {
	uid, _ := helpers_account.Auth(w, r)
	if uid < 0 {
		return nil, false
	}

	pidStr := r.PathValue("pid")
	nidStr := r.PathValue("nid")

	pid, perr := strconv.Atoi(pidStr)
	nid, nerr := strconv.Atoi(nidStr)
	if perr != nil || nerr != nil {
		http.Error(w, "Illegal project or note ID", http.StatusBadRequest)
		return nil, false
	}

	user := data.Users[uid]
	project := user.Projects[pid]

	return project.Notes[nid], true
}
