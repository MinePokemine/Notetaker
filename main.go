package main

import (
	"log"
	"net/http"

	handlers "github.com/MinePokemine/notetaker/Handlers"
	handlers_account "github.com/MinePokemine/notetaker/Handlers/Account"
	handlers_notes "github.com/MinePokemine/notetaker/Handlers/Notes"
	handlers_projects "github.com/MinePokemine/notetaker/Handlers/Projects"
	t "github.com/MinePokemine/notetaker/Helpers/Types"
	saveload "github.com/MinePokemine/notetaker/SaveLoad"
)

func main() {
	t.Users = saveload.Load("data.json")

	mux := http.NewServeMux()

	// Web App
	mux.HandleFunc("GET /", handlers.LoadHTML("Index/index.html"))

	mux.HandleFunc("GET /account/signup", handlers.LoadHTML("Account/signup.html"))
	mux.HandleFunc("GET /account/rename", handlers.LoadHTML("Account/rename.html"))
	mux.HandleFunc("GET /account", handlers.CreateFuncTemplateHandler("Account/my.html", handlers_account.MyAccount))

	mux.HandleFunc("GET /prjoects/new", handlers.LoadHTML("Projects/new.html"))
	mux.HandleFunc("GET /projects/{pid}", handlers.CreateFuncTemplateHandler("Projects/project.html", handlers_projects.Project))

	mux.HandleFunc("GET /projects/{pid}/newtag", handlers.CreateFuncTemplateHandler("Notes/newtag.html", handlers_projects.Project))
	mux.HandleFunc("GET /projects/{pid}/newnote", handlers.CreateFuncTemplateHandler("Notes/newnote.html", handlers_projects.Project))

	mux.HandleFunc("GET /projects/{pid}/search", handlers.CreateFuncTemplateHandler("Notes/startsearch.html", handlers_projects.Project))
	mux.HandleFunc("GET /projects/{pid}/searchpage", handlers.CreateFuncTemplateHandler("Notes/search.html", handlers_notes.Search))

	mux.HandleFunc("GET /projects/{pid}/notes/{nid}", handlers.CreateFuncTemplateHandler("Notes/note.html", handlers_notes.Note))
	mux.HandleFunc("GET /projects/{pid}/tags/{tid}", handlers.CreateFuncTemplateHandler("Notes/note.html", handlers_notes.Tag))

	mux.HandleFunc("GET /save", handlers.GetFuncRunHandler(func() {
		saveload.Save(t.Usrs(), "data.json")
	}, handlers.LoadHTML("saved.html")))

	// API
	mux.HandleFunc("POST /api/account/signup", handlers_account.CreateAccount)
	mux.HandleFunc("POST /api/account/rename", handlers_account.RenameAccount)

	mux.HandleFunc("POST /api/prjoects/new", handlers_projects.NewProject)

	mux.HandleFunc("POST /api/projects/{pid}/newtag", handlers_notes.NewTag)
	mux.HandleFunc("POST /api/projects/{pid}/newnote", handlers_notes.NewNote)

	port := ":8080"

	err := http.ListenAndServe(port, mux)

	saveload.Save(t.Users, "data.json")

	log.Fatalln(err.Error())
}
