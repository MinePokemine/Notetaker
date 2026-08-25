package handlers_notes

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	data "github.com/MinePokemine/notetaker/Data"
	helpers "github.com/MinePokemine/notetaker/Helpers"
	helpers_account "github.com/MinePokemine/notetaker/Helpers/Account"
)

func Search(w http.ResponseWriter, r *http.Request) ([]helpers.Note, bool) {
	fmt.Println("Searching")
	uid, _ := helpers_account.Auth(w, r)
	if uid < 0 {
		return []helpers.Note{}, false
	}
	user := data.Users[uid]

	pIDstr := r.PathValue("pid")
	pID, err := strconv.Atoi(pIDstr)
	if err != nil {
		http.Error(w, "Project ID not an int: "+err.Error(), http.StatusBadRequest)
		return []helpers.Note{}, false
	}
	if pID > len(user.Projects) {
		http.Error(w, "Invalid project "+pIDstr, http.StatusBadRequest)
		return []helpers.Note{}, false
	}
	project := user.Projects[pID]

	methodstr := r.FormValue("method")
	var isAny bool
	if methodstr == "" {
		http.Error(w, "Method not supplied", http.StatusBadRequest)
	}
	if strings.ToLower(methodstr) == "any" {
		isAny = true
	} else if strings.ToLower(methodstr) == "all" {
		isAny = false
	} else {
		http.Error(w, "Invalid method: "+methodstr+", should be either any or all", http.StatusBadRequest)
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return []helpers.Note{}, false
	}

	tagStrs := r.Form["tags"]

	var tags []*helpers.Tag

	for i, tagStr := range tagStrs {
		tID, err := strconv.Atoi(tagStr)
		if err != nil {
			http.Error(w, "Tag #"+strconv.Itoa(i)+" in the form not an int: "+err.Error(), http.StatusBadRequest)
			return []helpers.Note{}, false
		}
		if tID > len(project.Tags) {
			println("Invalid tag (Tag #" + strconv.Itoa(i) + " in form): " + strconv.Itoa(tID))
			continue
		}

		tags = append(tags, project.Tags[tID])
	}

	var notes []helpers.Note

	for _, note := range project.Notes {
		fmt.Println("Note")
		hasAny := false
		hasAll := true
		for _, tag := range tags {
			if !slices.ContainsFunc(note.Tags, tag.ChildOf) {
				hasAll = false
				fmt.Println("It's missing one")
			} else {
				hasAny = true
				fmt.Println("It has one")
			}
		}
		if isAny && hasAny {
			print("Any")
			notes = append(notes, *note)
		} else if !isAny && hasAll {
			print("All")
			notes = append(notes, *note)
		}
	}

	fmt.Println(notes)

	return notes, true
}
