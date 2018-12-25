package p

type P [3]int

func (p P) Dist(q P) (d int) {
	for i := range p {
		if b := p[i] - q[i]; b > 0 {
			d += b
		} else {
			d -= b
		}
	}
	return
}

type B struct {
	p P
	r int
}

func (b B) InRange(p P) bool {
	return b.p.Dist(p) <= b.r
}

func P1(input []B) (p1 int) {
	s := B{}
	for _, b := range input {
		if b.r > s.r {
			s = b
		}
	}
	for _, b := range input {
		if s.InRange(b.p) {
			p1++
		}
	}
	return
}
