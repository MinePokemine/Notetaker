package handlers

import (
	"fmt"
	"net/http"

	templates "github.com/MinePokemine/notetaker/Templates"
)

func CreateFuncTemplateHandler[T any](name string, getArgs func(http.ResponseWriter, *http.Request) (T, bool)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := templates.Template(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		args, work := getArgs(w, r)

		if !work {
			fmt.Println("Error Loading " + name)
			return
		}

		err = t.Execute(w, args)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
