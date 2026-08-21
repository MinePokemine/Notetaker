package main

import (
	"log"
	"net/http"

	handlers "github.com/MinePokemine/notetaker/Handlers"
	handlers_account "github.com/MinePokemine/notetaker/Handlers/Account"
	handlers_notes "github.com/MinePokemine/notetaker/Handlers/Notes"
	handlers_projects "github.com/MinePokemine/notetaker/Handlers/Projects"
)

func main() {
	// Web App
	http.HandleFunc("GET /", handlers.LoadHTML("Templates/index.html"))

	http.HandleFunc("GET /account/signup", handlers.LoadHTML("Templates/Account/signup.html"))
	http.HandleFunc("GET /account/rename", handlers.LoadHTML("Templates/Account/rename.html"))
	http.HandleFunc("GET /account", handlers.CreateFuncTemplateHandler("Templates/Account/my.html", handlers_account.MyAccount))

	http.HandleFunc("GET /prjoects/new", handlers.LoadHTML("Templates/Projects/new.html"))
	http.HandleFunc("GET /projects/{pid}", handlers.CreateFuncTemplateHandler("Templates/Projects/project.html", handlers_projects.Project))

	http.HandleFunc("GET /projects/{pid}/newtag", handlers.CreateFuncTemplateHandler("Templates/Notes/newtag.html", handlers_projects.Project))
	http.HandleFunc("GET /projects/{pid}/newnote", handlers.CreateFuncTemplateHandler("Templates/Notes/newnote.html", handlers_projects.Project))

	// API
	http.HandleFunc("POST /api/account/signup", handlers_account.CreateAccount)
	http.HandleFunc("POST /api/account/rename", handlers_account.RenameAccount)

	http.HandleFunc("POST /api/prjoects/new", handlers_projects.NewProject)

	http.HandleFunc("POST /api/projects/{pid}/newtag", handlers_notes.NewTag)
	http.HandleFunc("POST /api/projects/{pid}/newnote", handlers_notes.NewNote)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
