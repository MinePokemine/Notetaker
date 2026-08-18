package handlers

import (
	"html/template"
	"net/http"
)

func CreateFuncTemplateHandler[T any](path string, getArgs func(http.ResponseWriter, *http.Request) (T, bool)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		args, work := getArgs(w, r)

		if !work {
			return
		}

		err = t.Execute(w, args)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
