package handlers_notes

import (
	"net/http"
	"strconv"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func Tag(w http.ResponseWriter, r *http.Request) (*t.Tag, bool) {
	uid, _ := helpers_account.Auth(w, r)
	if uid < 0 {
		return nil, false
	}

	pidStr := r.PathValue("pid")
	tidStr := r.PathValue("tid")

	pid, perr := strconv.Atoi(pidStr)
	tid, terr := strconv.Atoi(tidStr)
	if perr != nil || terr != nil {
		http.Error(w, "Illegal project or tag ID", http.StatusBadRequest)
		return nil, false
	}

	user := t.Users[uid]
	project := user.Projects[pid]

	return project.Tags[tid], true
}
