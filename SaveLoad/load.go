package saveload

import (
	"encoding/json"
	"os"
	"path/filepath"

	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func Load(filename string) []*t.User {
	ex, err := os.Executable()
	if err != nil {
		panic("Error finding path of executable: " + err.Error())
	}
	exePath := filepath.Dir(ex)
	path := filepath.Join(exePath, filename)

	jsonned, err := os.ReadFile(path)
	if err != nil {
		//panic("Error loading save file: " + err.Error())
		return make([]*t.User, 0)
	}

	users := make([]*t.User, 0)
	err = json.Unmarshal(jsonned, users)
	if err != nil {
		panic("Error reading save file json: " + err.Error())
	}

	return users
}
