package t

func (a *Tag) ChildOf(b *Tag) bool {
	if a == b {
		return true
	}

	bottomUp := make(Set[*Tag])
	bottomUnchecked := make(Set[*Tag])
	bottomUnchecked.Add(a)

	topDown := make(Set[*Tag])
	topUnchecked := make(Set[*Tag])
	topUnchecked.Add(b)

	for len(bottomUnchecked) > 0 && len(topUnchecked) > 0 {
		bottomUp.AddAll(bottomUnchecked)
		topDown.AddAll(topUnchecked)

		newB := make(Set[*Tag])
		for b := range bottomUnchecked {
			if topDown.Contains(b) {
				return true
			}
			newB.AddAllSlice(b.Parents)
		}
		bottomUnchecked = newB

		newT := make(Set[*Tag])
		for t := range topUnchecked {
			if bottomUp.Contains(t) {
				return true
			}
			newT.AddAllSlice(t.Children)
		}
		topUnchecked = newT
	}

	return false
}
