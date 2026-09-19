package saveload

import (
	"encoding/json"
	"os"
	"path/filepath"

	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func Save(data []*t.User, filename string) {
	jsonned, err := json.Marshal(data)
	if err != nil {
		panic("Error converting data to JSON: " + err.Error())
	}

	ex, err := os.Executable()
	if err != nil {
		panic("Error finding path of executable: " + err.Error())
	}
	exePath := filepath.Dir(ex)
	path := filepath.Join(exePath, filename)

	f, err := os.Create(path)
	if err != nil {
		panic("Error creating save file: " + err.Error())
	}

	_, err = f.Write(jsonned)
	if err != nil {
		panic("Error writing to save file: " + err.Error())
	}
}
