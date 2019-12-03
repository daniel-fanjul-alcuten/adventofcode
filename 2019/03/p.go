package p

type point struct {
	x, y int
}

func (p point) distance() (d int) {
	if p.x >= 0 {
		d += p.x
	} else {
		d -= p.x
	}
	if p.y >= 0 {
		d += p.y
	} else {
		d -= p.y
	}
	return
}

type direction rune

func (d direction) next(p point) point {
	switch d {
	case 'D':
		return point{p.x, p.y - 1}
	case 'L':
		return point{p.x - 1, p.y}
	case 'R':
		return point{p.x + 1, p.y}
	case 'U':
		return point{p.x, p.y + 1}
	}
	panic(d)
}

type step struct {
	d direction
	n int
}

type wire map[point]int

func (w *wire) fill(p point, ss ...step) {
	d := 0
	*w = wire{p: d}
	for _, s := range ss {
		for i := 0; i < s.n; i++ {
			p, d = s.d.next(p), d+1
			if _, ok := (*w)[p]; !ok {
				(*w)[p] = d
			}
		}
	}
	return
}

func (w *wire) intersect(w1, w2 wire) {
	*w = wire{}
	for p, d1 := range w1 {
		if d2, ok := w2[p]; ok {
			(*w)[p] = d1 + d2
		}
	}
	return
}

func (w wire) distance1() (d int) {
	for p := range w {
		pd := p.distance()
		if d == 0 || (pd > 0 && pd < d) {
			d = pd
		}
	}
	return
}

func (w wire) distance2() (d int) {
	for _, pd := range w {
		if d == 0 || (pd > 0 && pd < d) {
			d = pd
		}
	}
	return
}

func P1(s1, s2 []step) int {
	var w1, w2, w3 wire
	w1.fill(point{0, 0}, s1...)
	w2.fill(point{0, 0}, s2...)
	w3.intersect(w1, w2)
	return w3.distance1()
}

func P2(s1, s2 []step) int {
	var w1, w2, w3 wire
	w1.fill(point{0, 0}, s1...)
	w2.fill(point{0, 0}, s2...)
	w3.intersect(w1, w2)
	return w3.distance2()
}
