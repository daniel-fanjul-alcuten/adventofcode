package p

import (
	"strconv"
	"strings"
)

func P1(input []string, me bool) (output int) {
	type P struct {
		S, T string
	}
	people := map[string]struct{}{}
	happiness := map[P]int{}
	for _, in := range input {
		s := strings.Split(in, " ")
		ps := s[0]
		pt := strings.Trim(s[10], ".")
		h, err := strconv.Atoi(s[3])
		if err != nil {
			panic(err)
		}
		people[ps] = struct{}{}
		people[pt] = struct{}{}
		if s[2] == "gain" {
			happiness[P{ps, pt}] = h
		} else {
			happiness[P{ps, pt}] = -h
		}
	}
	if me {
		people["me"] = struct{}{}
	}
	var find func(f, p string) int
	find = func(f, p string) int {
		if len(people) == 1 {
			for q := range people {
				return happiness[P{p, q}] + happiness[P{q, p}] +
					happiness[P{q, f}] + happiness[P{f, q}]
			}
		}
		pp := make([]string, 0, len(people))
		for p := range people {
			pp = append(pp, p)
		}
		best := 0
		for _, q := range pp {
			delete(people, q)
			h := happiness[P{p, q}] + happiness[P{q, p}] + find(f, q)
			people[q] = struct{}{}
			if best == 0 || h > best {
				best = h
			}
		}
		return best
	}
	pp := make([]string, 0, len(people))
	for p := range people {
		pp = append(pp, p)
	}
	best := 0
	for _, p := range pp {
		delete(people, p)
		h := find(p, p)
		people[p] = struct{}{}
		if best == 0 || h > best {
			best = h
		}
	}
	return best
}
