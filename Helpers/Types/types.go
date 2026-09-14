package t

type (
	Note struct {
		IDInProject int
		Project     *Project
		Data        string
		Source      string
		Tags        []*Tag
	}

	Tag struct {
		IDInProject int
		Project     *Project
		Name        string
		Children    []*Tag
		Parents     []*Tag
		Notes       []*Note
	}

	Project struct {
		Name     string
		User     *User
		IDInUser int
		Notes    []*Note
		Tags     []*Tag
	}

	User struct {
		Username string
		UID      int
		Login    [32]byte
		Projects []*Project
	}
)
