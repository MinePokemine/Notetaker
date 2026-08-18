package helpers

type (
	Note struct {
		Value string
		Tags  []*Tag
	}

	Tag struct {
		Name     string
		Children []*Tag
		Parents  []*Tag
		Notes    []*Note
	}

	Project struct {
		Name  string
		User  *User
		Notes []*Note
		Tags  []*Tag
	}

	User struct {
		Username string
		UID      int
		Login    [32]byte
		Projects []*Project
	}
)
