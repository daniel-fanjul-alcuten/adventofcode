package p

import (
	"sort"
	"strconv"
	"strings"
)

type quantity struct {
	n int
	m string
}

type reaction struct {
	i []quantity
	o quantity
}

type reactions map[string]reaction

func parse(ss []string) (m reactions) {
	m = make(reactions, len(ss))
	for _, s := range ss {
		r := reaction{}
		sp1 := strings.Split(s, " => ")
		sp2 := strings.Split(sp1[0], ", ")
		sp := func(s string) quantity {
			sp3 := strings.Split(s, " ")
			n, err := strconv.Atoi(sp3[0])
			if err != nil {
				panic(err)
			}
			return quantity{n, sp3[1]}
		}
		for _, s := range sp2 {
			r.i = append(r.i, sp(s))
		}
		r.o = sp(sp1[1])
		m[r.o.m] = r
	}
	return
}

func P1(fuel int, rr reactions) (ore int) {
	pending := map[string]int{"FUEL": fuel}
	for {
		runs := 0
		for m, q := range pending {
			if q <= 0 {
				continue
			}
			r := rr[m]
			n := q / r.o.n
			if n*r.o.n < q {
				n++
			}
			pending[m] -= n * r.o.n
			if pending[m] == 0 {
				delete(pending, m)
			}
			for _, q := range r.i {
				if q.m == "ORE" {
					ore += n * q.n
				} else {
					pending[q.m] += n * q.n
				}
			}
			runs++
		}
		if runs == 0 {
			break
		}
	}
	return
}

func P2(rr reactions) (fuel int) {
	fuel = sort.Search(1000000000000, func(fuel int) bool {
		ore := P1(fuel, rr)
		if ore < 1000000000000 {
			return false
		} else {
			return true
		}
	})
	for P1(fuel, rr) > 1000000000000 {
		fuel--
	}
	return
}
