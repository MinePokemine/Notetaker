package handlers

import (
	"html/template"
	"net/http"
)

func CreateTemplateHandler[T any](path string, args T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.Execute(w, args)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
