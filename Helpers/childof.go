package helpers

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
