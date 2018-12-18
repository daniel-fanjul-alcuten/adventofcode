package p

import (
	"strconv"
	"strings"
)

func P1(input []string, r bool) int {
	type P struct {
		S, T string
	}
	l := map[string]struct{}{}
	d := map[P]int{}
	for _, in := range input {
		s := strings.Split(in, " ")
		ls, lt := s[0], s[2]
		m, err := strconv.Atoi(s[4])
		if err != nil {
			panic(err)
		}
		l[ls] = struct{}{}
		l[lt] = struct{}{}
		d[P{ls, lt}] = m
		d[P{lt, ls}] = m
	}
	var rec func(c string) int
	rec = func(c string) int {
		if len(l) == 1 {
			for n := range l {
				return d[P{c, n}]
			}
		}
		ll := make([]string, 0, len(l))
		for n := range l {
			ll = append(ll, n)
		}
		best := 0
		for _, n := range ll {
			delete(l, n)
			dn := d[P{c, n}] + rec(n)
			l[n] = struct{}{}
			if best == 0 || !r && dn < best || r && dn > best {
				best = dn
			}
		}
		return best
	}
	ll := make([]string, 0, len(l))
	for n := range l {
		ll = append(ll, n)
	}
	best := 0
	for _, n := range ll {
		delete(l, n)
		dn := rec(n)
		l[n] = struct{}{}
		if best == 0 || !r && dn < best || r && dn > best {
			best = dn
		}
	}
	return best
}
