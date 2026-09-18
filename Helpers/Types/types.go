package t

type (
	Note struct {
		UID int
		PID int
		NID int

		Data   string
		Source string
		Tags   []*Tag
	}

	Tag struct {
		UID int
		PID int
		TID int

		Name     string
		Children []*Tag
		Parents  []*Tag
		Notes    []*Note
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
