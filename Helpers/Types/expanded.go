package t

type (
	NoteEx struct {
		Project ProjectEx
		NID     int

		Data   string
		Source string
		Tags   []TagEx
	}

	TagEx struct {
		Project ProjectEx
		TID     int

		Name     string
		Children []TagEx
		Parents  []TagEx
		Notes    []NoteEx
	}

	ProjectEx struct {
		User User
		PID  int

		Name  string
		Notes []NoteEx
		Tags  []TagEx
	}
)
