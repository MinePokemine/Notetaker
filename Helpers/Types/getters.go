package t

func (n Note) User() *User {
	return Users[n.UID]
}
func (n Note) Project() *Project {
	return n.User().Projects[n.PID]
}

func (t Tag) User() *User {
	return Users[t.UID]
}
func (t Tag) Project() *Project {
	return t.User().Projects[t.PID]
}

func (p Project) User() *User {
	return Users[p.UID]
}

func (tr TagReference) Get() *Tag {
	return Users[tr.UID].Projects[tr.PID].Tags[tr.TID]
}
func (nr NoteReference) Get() *Note {
	return Users[nr.UID].Projects[nr.PID].Notes[nr.NID]
}

func GetTags(trs []TagReference) []*Tag {
	var tags []*Tag
	for _, tr := range trs {
		tags = append(tags, tr.Get())
	}
	return tags
}

func GetNotes(nrs []NoteReference) []*Note {
	var notes []*Note
	for _, nr := range nrs {
		notes = append(notes, nr.Get())
	}
	return notes
}

func (t *Tag) Reference() TagReference {
	return TagReference{
		UID: t.UID,
		PID: t.PID,
		TID: t.TID,
	}
}

func (n *Note) Reference() NoteReference {
	return NoteReference{
		UID: n.UID,
		PID: n.PID,
		NID: n.NID,
	}
}
