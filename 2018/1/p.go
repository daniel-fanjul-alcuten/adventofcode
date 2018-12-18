package p

func P1(s int, xx []int) (r int) {
	r = s
	for _, x := range xx {
		r += x
	}
	return
}

func P2(s int, xx []int) (r int) {
	m := map[int]struct{}{}
	r = s
	m[r] = struct{}{}
	for {
		for _, x := range xx {
			r += x
			if _, ok := m[r]; ok {
				return
			}
			m[r] = struct{}{}
		}
	}
}
