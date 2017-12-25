package p24

import (
	"strconv"
	"strings"
)

type Components map[int]map[int]int

func (c Components) Add(a, b int) {
	m := c[a]
	if m == nil {
		m = map[int]int{}
		c[a] = m
	}
	m[b] += 1
}

func Parse(ss []string) (c Components) {
	c = Components{}
	for _, s := range ss {
		p := strings.Split(s, "/")
		a, err := strconv.Atoi(p[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(p[1])
		if err != nil {
			panic(err)
		}
		c.Add(a, b)
		c.Add(b, a)
	}
	return
}

func Strongest(p, s int, r Components) (ss int) {
	ss = s
	for q, n := range r[p] {
		if n > 0 {
			r[p][q] -= 1
			r[q][p] -= 1
			if s2 := Strongest(q, s+p+q, r); s2 > ss {
				ss = s2
			}
			r[p][q] += 1
			r[q][p] += 1
		}
	}
	return
}

func LongestStrongest(p, l, s int, r Components) (ll, ss int) {
	ll, ss = l, s
	for q, n := range r[p] {
		if n > 0 {
			r[p][q] -= 1
			r[q][p] -= 1
			if l2, s2 := LongestStrongest(q, l+1, s+p+q, r); l2 > ll || (l2 == ll && s2 > ss) {
				ll, ss = l2, s2
			}
			r[p][q] += 1
			r[q][p] += 1
		}
	}
	return
}
