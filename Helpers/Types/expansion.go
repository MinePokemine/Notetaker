package t

// Notes
func (n Note) Expand() NoteEx {
	return NoteEx{
		Project: n.Project().Expand(),
		NID:     n.NID,

		Data:   n.Data,
		Source: n.Source,
		Tags:   ExpandTRs(n.Tags),
	}
}

func (nr NoteReference) Expand() NoteEx {
	return nr.Get().Expand()
}

func ExpandNRs(nrs []NoteReference) []NoteEx {
	nex := make([]NoteEx, 0)
	for _, nr := range nrs {
		nex = append(nex, nr.Expand())
	}
	return nex
}

func ExpandNotes(ns []*Note) []NoteEx {
	nex := make([]NoteEx, 0)
	for _, n := range ns {
		nex = append(nex, n.Expand())
	}
	return nex
}

// Tags
func (t Tag) Expand() TagEx {
	return TagEx{
		Project: t.Project().Expand(),
		TID:     t.TID,

		Name:     t.Name,
		Children: ExpandTRs(t.Children),
		Parents:  ExpandTRs(t.Parents),
		Notes:    ExpandNRs(t.Notes),
	}
}

func (tr TagReference) Expand() TagEx {
	return tr.Get().Expand()
}

func ExpandTRs(trs []TagReference) []TagEx {
	tex := make([]TagEx, 0)
	for _, tr := range trs {
		tex = append(tex, tr.Expand())
	}
	return tex
}

func ExpandTags(ts []*Tag) []TagEx {
	tex := make([]TagEx, 0)
	for _, t := range ts {
		tex = append(tex, t.Expand())
	}
	return tex
}

// Projects
func (p Project) Expand() ProjectEx {
	return ProjectEx{
		User: *Users[p.UID],
		PID:  p.PID,

		Name:  p.Name,
		Notes: ExpandNotes(p.Notes),
	}
}
