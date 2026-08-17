package handlers

import (
	"html/template"
	"net/http"
)

func CreateFuncTemplateHandler[T any](path string, getArgs func(http.ResponseWriter, *http.Request) T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.Execute(w, getArgs(w, r))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
