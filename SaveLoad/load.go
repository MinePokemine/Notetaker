package saveload

import (
	"encoding/json"
	"os"

	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func Load() []*t.User {
	jsonned, err := os.ReadFile(SAVEPATH)
	if err != nil {
		//panic("Error loading save file: " + err.Error())
		return make([]*t.User, 0)
	}

	usersNP := make([]*t.User, 0)
	users := &usersNP
	err = json.Unmarshal(jsonned, users)
	if err != nil {
		panic("Error reading save file json: " + err.Error())
	}

	return *users
}
