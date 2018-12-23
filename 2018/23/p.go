package p

type P struct {
	x, y, z int
}

func (p P) Dist(q P) (d int) {
	if p.x >= q.x {
		d += p.x - q.x
	} else {
		d += q.x - p.x
	}
	if p.y >= q.y {
		d += p.y - q.y
	} else {
		d += q.y - p.y
	}
	if p.z >= q.z {
		d += p.z - q.z
	} else {
		d += q.z - p.z
	}
	return
}

type Bot struct {
	p P
	r int
}

func P1(input []Bot) (p1 int) {
	s := Bot{}
	for _, b := range input {
		if b.r > s.r {
			s = b
		}
	}
	for _, b := range input {
		if s.p.Dist(b.p) <= s.r {
			p1++
		}
	}
	return
}
