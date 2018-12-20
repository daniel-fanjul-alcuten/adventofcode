package p

import (
	"strings"
)

func P1(regex string) (n, r int) {
	type P struct {
		x, y int
	}
	adjs := func(p P) []P {
		return []P{
			P{p.x - 1, p.y},
			P{p.x + 1, p.y},
			P{p.x, p.y - 1},
			P{p.x, p.y + 1},
		}
	}
	type D struct {
		p1, p2 P
	}
	doors := map[D]struct{}{}
	regex = strings.TrimLeft(regex, "^")
	regex = strings.TrimRight(regex, "$")
	ppp := []map[P]struct{}{map[P]struct{}{P{}: struct{}{}}}
	qqq := []map[P]struct{}{}
	for i := 0; i < len(regex); i++ {
		switch regex[i] {
		case '(':
			pp := ppp[len(ppp)-1]
			pp2 := make(map[P]struct{}, len(pp))
			for p := range pp {
				pp2[p] = struct{}{}
			}
			ppp = append(ppp, pp2)
			qqq = append(qqq, map[P]struct{}{})
		case '|':
			pp := ppp[len(ppp)-1]
			qq := qqq[len(qqq)-1]
			for p := range pp {
				delete(pp, p)
				qq[p] = struct{}{}
			}
			pp0 := ppp[len(ppp)-2]
			for p := range pp0 {
				pp[p] = struct{}{}
			}
		case ')':
			pp := ppp[len(ppp)-1]
			qq := qqq[len(qqq)-1]
			for p := range pp {
				delete(pp, p)
				qq[p] = struct{}{}
			}
			ppp[len(ppp)-1], ppp = nil, ppp[:len(ppp)-1]
			qqq[len(qqq)-1], qqq = nil, qqq[:len(qqq)-1]
			ppp[len(ppp)-1] = qq
		default:
			x, y := 0, 0
			switch regex[i] {
			case 'N':
				y--
			case 'E':
				x++
			case 'W':
				x--
			case 'S':
				y++
			default:
				panic(regex[i : i+1])
			}
			pp := ppp[len(ppp)-1]
			pp2 := make(map[P]struct{}, len(pp))
			for p := range pp {
				q := P{p.x + x, p.y + y}
				doors[D{p, q}] = struct{}{}
				doors[D{q, p}] = struct{}{}
				pp2[q] = struct{}{}
			}
			ppp[len(ppp)-1] = pp2
		}
	}
	type S struct {
		p P
		n int
	}
	pending, processed := []S{S{P{}, 0}}, map[P]struct{}{P{}: struct{}{}}
	for len(pending) > 0 {
		s := pending[0]
		pending = pending[1:]
		for _, p := range adjs(s.p) {
			if _, ok := processed[p]; ok {
				continue
			}
			if _, ok := doors[D{s.p, p}]; !ok {
				continue
			}
			nn := s.n + 1
			if nn > n {
				n = nn
			}
			if nn >= 1000 {
				r++
			}
			pending = append(pending, S{p, nn})
			processed[p] = struct{}{}
		}
	}
	return
}
