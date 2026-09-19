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
		UID int `json:"uid"`
		PID int `json:"pid"`
		TID int `json:"tid"`

		Name     string          `json:"name"`
		Children []TagReference  `json:"children"`
		Parents  []TagReference  `json:"parents"`
		Notes    []NoteReference `json:"notes"`
	}

	Project struct {
		UID int `json:"uid"`
		PID int `json:"pid"`

		Name  string  `json:"name"`
		Notes []*Note `json:"notes"`
		Tags  []*Tag  `json:"tags"`
	}

	User struct {
		UID int `json:"uid"`

		Username string     `json:"name"`
		Login    [32]byte   `json:"login"`
		Projects []*Project `json:"projs"`
	}
)
