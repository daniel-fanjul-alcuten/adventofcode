package p

func P1A(s string) (l2, l3 bool) {
	m1, m2, m3 := map[rune]int{}, map[rune]struct{}{}, map[rune]struct{}{}
	for _, r := range s {
		m1[r]++
		switch {
		case m1[r] == 2:
			m2[r] = struct{}{}
		case m1[r] == 3:
			delete(m2, r)
			m3[r] = struct{}{}
		case m1[r] > 3:
			delete(m3, r)
		}
	}
	l2, l3 = len(m2) > 0, len(m3) > 0
	return
}

func P1B(ss []string) (k int) {
	c2, c3 := 0, 0
	for _, s := range ss {
		l2, l3 := P1A(s)
		if l2 {
			c2++
		}
		if l3 {
			c3++
		}
	}
	k = c2 * c3
	return
}

func P2A(s1, s2 string) (s3 string, ok bool) {
	i, d := 0, 0
	for i < len(s1) {
		if s1[i] == s2[i] {
			s3 = s3 + string(s1[i])
			i++
			continue
		}
		if d < 1 {
			d++
			i++
			continue
		}
		s3, ok = "", false
		return
	}
	ok = true
	return
}

func P2B(ss []string) (tt []string) {
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			if t, ok := P2A(ss[i], ss[j]); ok {
				tt = append(tt, t)
			}
		}
	}
	return
}
