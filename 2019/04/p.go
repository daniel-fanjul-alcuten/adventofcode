package p

import "strconv"

func P1(min, max int) (n int) {
	for v := min; v <= max; v++ {
		s := strconv.Itoa(v)
		rep := false
		for i := 1; i < len(s); i++ {
			if s[i-1] > s[i] {
				rep = true
				break
			}
		}
		if rep {
			continue
		}
		adj := false
		for i := 1; i < len(s); i++ {
			if s[i-1] == s[i] {
				adj = true
				break
			}
		}
		if !adj {
			continue
		}
		n++
	}
	return
}

func P2(min, max int) (n int) {
	for v := min; v <= max; v++ {
		s := strconv.Itoa(v)
		rep := false
		for i := 1; i < len(s); i++ {
			if s[i-1] > s[i] {
				rep = true
				break
			}
		}
		if rep {
			continue
		}
		nadj, adj := 1, false
		for i := 1; i < len(s); i++ {
			if s[i-1] == s[i] {
				nadj++
			} else {
				if nadj == 2 {
					adj = true
				}
				nadj = 1
			}
		}
		if nadj == 2 {
			adj = true
		}
		if !adj {
			continue
		}
		n++
	}
	return
}
