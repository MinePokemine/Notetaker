package helpers

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

func (a *Tag) ChildOf(b *Tag) bool {
	if a == b {
		return true
	}

	is := false
	for _, t := range a.Parents {
		if t.ChildOf(b) {
			is = true
		}
	}

	return is
}
