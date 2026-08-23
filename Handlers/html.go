package handlers

import (
	"net/http"

	templates "github.com/MinePokemine/notetaker/Templates"
)

func LoadHTML(name string) func(http.ResponseWriter, *http.Request) {
	// Read the entire file into memory
	content, err := templates.ReadFile(name)
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Failed to load html file "+name, http.StatusInternalServerError)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(content)
	}
}
