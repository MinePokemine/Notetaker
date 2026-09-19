package t

func (a *Tag) ChildOfTag(b *Tag) bool {
	return a.Reference().ChildOfTag(b)
}

func (a *Tag) ChildOfRef(b TagReference) bool {
	return a.Reference().ChildOfRef(b)
}

func (a TagReference) ChildOfTag(b *Tag) bool {
	return a.ChildOfRef(b.Reference())
}

func (a TagReference) ChildOfRef(b TagReference) bool {
	if a == b {
		return true
	}

	bottomUp := make(Set[TagReference])
	bottomUnchecked := make(Set[TagReference])
	bottomUnchecked.Add(a)

	topDown := make(Set[TagReference])
	topUnchecked := make(Set[TagReference])
	topUnchecked.Add(b)

	for len(bottomUnchecked) > 0 && len(topUnchecked) > 0 {
		bottomUp.AddAll(bottomUnchecked)
		topDown.AddAll(topUnchecked)

		newB := make(Set[TagReference])
		for b := range bottomUnchecked {
			if bottomUp.Contains(b) {
				continue
			}
			bottomUp.Add(b)
			if topDown.Contains(b) || topUnchecked.Contains(b) {
				return true
			}
			newB.AddAllSlice(b.Get().Parents)
		}
		bottomUnchecked = newB

		newT := make(Set[TagReference])
		for t := range topUnchecked {
			if topDown.Contains(t) {
				continue
			}
			topDown.Add(t)
			if bottomUp.Contains(t) || bottomUnchecked.Contains(t) {
				return true
			}
			newT.AddAllSlice(t.Get().Children)
		}
		topUnchecked = newT
	}

	return false
}
