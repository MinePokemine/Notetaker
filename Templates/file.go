package templates

import (
	"embed"
	"html/template"
	"io/fs"
)

//go:embed all:*
var Files embed.FS

func ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(Files, name)
}

func Template(name string) (*template.Template, error) {
	return template.ParseFS(Files, name)
}
