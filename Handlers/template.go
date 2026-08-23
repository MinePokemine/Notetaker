package handlers

import (
	"net/http"

	templates "github.com/MinePokemine/notetaker/Templates"
)

func CreateTemplateHandler[T any](name string, args T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := templates.Template(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, args)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
