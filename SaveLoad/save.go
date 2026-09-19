package saveload

import (
	"encoding/json"
	"os"

	t "github.com/MinePokemine/notetaker/Helpers/Types"
)

func Save(data []*t.User) {
	jsonned, err := json.Marshal(data)
	if err != nil {
		panic("Error converting data to JSON: " + err.Error())
	}

	f, err := os.Create(SAVEPATH)
	if err != nil {
		panic("Error creating save file: " + err.Error())
	}

	_, err = f.Write(jsonned)
	if err != nil {
		panic("Error writing to save file: " + err.Error())
	}
}
