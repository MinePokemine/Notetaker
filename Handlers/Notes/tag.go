package handlers_notes

import (
	"net/http"
	"strconv"

	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func Tag(w http.ResponseWriter, r *http.Request) (t.TagEx, bool) {
	uid, _ := helpers_account.Auth(w, r)
	if uid < 0 {
		return t.TagEx{}, false
	}

	pidStr := r.PathValue("pid")
	tidStr := r.PathValue("tid")

	pid, perr := strconv.Atoi(pidStr)
	tid, terr := strconv.Atoi(tidStr)
	if perr != nil || terr != nil {
		http.Error(w, "Illegal project or tag ID", http.StatusBadRequest)
		return t.TagEx{}, false
	}

	user := t.Users[uid]
	project := user.Projects[pid]

	return project.Tags[tid].Expand(), true
}
