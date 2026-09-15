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
			if bottomUp.Contains(b) {
				continue
			}
			bottomUp.Add(b)
			if topDown.Contains(b) || topUnchecked.Contains(b) {
				return true
			}
			newB.AddAllSlice(b.Parents)
		}
		bottomUnchecked = newB

		newT := make(Set[*Tag])
		for t := range topUnchecked {
			if topDown.Contains(t) {
				continue
			}
			topDown.Add(t)
			if bottomUp.Contains(t) || bottomUnchecked.Contains(t) {
				return true
			}
			newT.AddAllSlice(t.Children)
		}
		topUnchecked = newT
	}

	return false
}
