package saveload

import (
	"os"
	"path/filepath"
)

var SAVEPATH = "/users/rakesivar/Documents/GoLang/Notetaker/data.json"

func GeneratePath(filename string) {
	ex, err := os.Executable()
	if err != nil {
		panic("Error finding path of executable: " + err.Error())
	}
	exePath := filepath.Dir(ex)
	SAVEPATH = filepath.Join(exePath, filename)
}
