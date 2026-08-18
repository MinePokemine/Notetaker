package handlers

import (
	"net/http"
	"os"
)

func LoadHTML(path string) func(http.ResponseWriter, *http.Request) {
	// Read the entire file into memory
	content, err := os.ReadFile(path)
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Failed to load html file "+path, http.StatusInternalServerError)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(content)
	}
}
