package t

type (
	NoteReference struct {
		UID int `json:"uid"`
		PID int `json:"pid"`
		NID int `json:"nid"`
	}

	Note struct {
		UID int `json:"uid"`
		PID int `json:"pid"`
		NID int `json:"nid"`

		Data   string         `json:"data"`
		Source string         `json:"src"`
		Tags   []TagReference `json:"tags"`
	}

	TagReference struct {
		UID int `json:"uid"`
		PID int `json:"pid"`
		TID int `json:"tid"`
	}

	Tag struct {
		UID int
		PID int
		TID int

		Name     string
		Children []TagReference
		Parents  []TagReference
		Notes    []NoteReference
	}

	Project struct {
		UID int
		PID int

		Name  string
		Notes []*Note
		Tags  []*Tag
	}

	User struct {
		UID int

		Username string
		Login    [32]byte
		Projects []*Project
	}
)
