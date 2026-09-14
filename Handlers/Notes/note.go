package handlers_notes

import (
	"net/http"
	"strconv"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func Note(w http.ResponseWriter, r *http.Request) (*t.Note, bool) {
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

	user := t.Users[uid]
	project := user.Projects[pid]

	return project.Notes[nid], true
}
