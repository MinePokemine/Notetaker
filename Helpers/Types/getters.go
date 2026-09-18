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
