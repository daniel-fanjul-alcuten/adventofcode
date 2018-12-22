package p

type P struct {
	x, y int
}

func P12(mouth, target P, depth int) (risk, mins int) {
	geoindex := map[P]int{}
	geoindex[mouth] = 0
	geoindex[target] = 0
	var erolevelf func(p P) int
	geoindexf := func(p P) int {
		if g, ok := geoindex[p]; ok {
			return g
		}
		if p.y == 0 {
			geoindex[p] = 16807 * p.x
			return geoindex[p]
		}
		if p.x == 0 {
			geoindex[p] = 48271 * p.y
			return geoindex[p]
		}
		geoindex[p] = erolevelf(P{p.x - 1, p.y}) * erolevelf(P{p.x, p.y - 1})
		return geoindex[p]
	}
	erolevelf = func(p P) int {
		return (geoindexf(p) + depth) % 20183
	}
	regtype := func(p P) int {
		return erolevelf(p) % 3
	}
	for y := mouth.y; y <= target.y; y++ {
		for x := mouth.x; x <= target.x; x++ {
			risk += regtype(P{x, y})
		}
	}
	type S struct {
		p P
		t string
	}
	time, states := 0, [][]S{[]S{S{mouth, "t"}}}
	add := func(i int, s S) {
		for i >= len(states) {
			states = append(states, []S{})
		}
		states[i] = append(states[i], s)
	}
	processed := map[S]struct{}{states[0][0]: struct{}{}}
	for len(states) > 0 {
		if len(states[0]) == 0 {
			time++
			states = states[1:]
			continue
		}
		s := states[0][0]
		states[0] = states[0][1:]
		if s == (S{target, "t"}) {
			mins = time
			return
		}
		for _, p2 := range []P{
			P{s.p.x - 1, s.p.y},
			P{s.p.x + 1, s.p.y},
			P{s.p.x, s.p.y - 1},
			P{s.p.x, s.p.y + 1},
		} {
			if p2.x < 0 {
				continue
			}
			if p2.y < 0 {
				continue
			}
			s2 := S{p2, s.t}
			if _, ok := processed[s2]; ok {
				continue
			}
			processed[s2] = struct{}{}
			good := false
			switch pt2 := regtype(p2); pt2 {
			case 0:
				good = s.t == "c" || s.t == "t"
			case 1:
				good = s.t == "c" || s.t == "n"
			case 2:
				good = s.t == "t" || s.t == "n"
			default:
				panic(pt2)
			}
			if good {
				add(1, s2)
			}
		}
		for _, t2 := range []string{"t", "c", "n"} {
			if s.t == t2 {
				continue
			}
			s2 := S{s.p, t2}
			if _, ok := processed[s2]; ok {
				continue
			}
			good := false
			switch pt := regtype(s.p); pt {
			case 0:
				good = t2 == "c" || t2 == "t"
			case 1:
				good = t2 == "c" || t2 == "n"
			case 2:
				good = t2 == "t" || t2 == "n"
			default:
				panic(pt)
			}
			if good {
				add(7, s2)
			}
		}
	}
	return
}
