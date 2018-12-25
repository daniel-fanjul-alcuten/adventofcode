package p

type P [4]int

func (p P) Dist(q P) (d int) {
	for i := range p {
		if b := p[i] - q[i]; b >= 0 {
			d += b
		} else {
			d -= b
		}
	}
	return
}

func (p P) InRangeP(q P) bool {
	return p.Dist(q) <= 3
}

func (p P) InRangeS(s []P) bool {
	for _, q := range s {
		if p.InRangeP(q) {
			return true
		}
	}
	return false
}

func (p P) InRangeSS(ss [][]P) (is []int) {
	for i, s := range ss {
		if p.InRangeS(s) {
			is = append(is, i)
		}
	}
	return
}

func P1(input []P) (n int) {
	ss := [][]P{}
	for _, p := range input {
		is := p.InRangeSS(ss)
		switch len(is) {
		case 0:
			ss = append(ss, []P{p})
		case 1:
			i := is[0]
			ss[i] = append(ss[i], p)
		default:
			ss2 := [][]P{[]P{p}}
			for i := range ss {
				if len(is) > 0 && i == is[0] {
					ss2[0] = append(ss2[0], ss[i]...)
					is = is[1:]
					continue
				}
				ss2 = append(ss2, ss[i])
				continue
			}
			ss = ss2
		}
	}
	return len(ss)
}
