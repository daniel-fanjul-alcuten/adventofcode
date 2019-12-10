package p

import "sort"

type point struct {
	x, y int
}

func (p point) sub(q point) point {
	return point{p.x - q.x, p.y - q.y}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (p point) div() (d int, q point) {
	x := p.x
	if x < 0 {
		x = -x
	}
	y := p.y
	if y < 0 {
		y = -y
	}
	d = gcd(x, y)
	q = point{p.x / d, p.y / d}
	return
}

func parse(ss ...string) (pp []point) {
	for y, s := range ss {
		for x, r := range s {
			if r != '.' {
				pp = append(pp, point{x, y})
			}
		}
	}
	return
}

func (p point) visible(qq []point) (v int) {
	m := map[point]struct{}{}
	for _, q := range qq {
		if p == q {
			continue
		}
		_, g := q.sub(p).div()
		m[g] = struct{}{}
	}
	v = len(m)
	return
}

func P1(pp []point) (q point, n int) {
	for _, p := range pp {
		if v := p.visible(pp); v > n {
			q, n = p, v
		}
	}
	return
}

func (p point) add(q point) point {
	return point{p.x + q.x, p.y + q.y}
}

func (p point) mul(d int) point {
	return point{p.x * d, p.y * d}
}

type vector struct {
	g  point
	dd []int
}

func (v vector) quad() int {
	if v.g.x == 0 && v.g.y < 0 {
		return 0
	} else if v.g.x > 0 && v.g.y < 0 {
		return 1
	} else if v.g.x > 0 && v.g.y == 0 {
		return 2
	} else if v.g.x > 0 && v.g.y > 0 {
		return 3
	} else if v.g.x == 0 && v.g.y > 0 {
		return 4
	} else if v.g.x < 0 && v.g.y > 0 {
		return 5
	} else if v.g.x < 0 && v.g.y == 0 {
		return 6
	} else if v.g.x < 0 && v.g.y < 0 {
		return 7
	}
	panic(8)
}

type vectors []vector

func (vv *vectors) Len() int {
	return len(*vv)
}

func (vv *vectors) Less(i, j int) bool {
	vi := (*vv)[i]
	vj := (*vv)[j]
	qi := vi.quad()
	qj := vj.quad()
	if qi < qj {
		return true
	} else if qi > qj {
		return false
	}
	gi := vi.g
	gj := vj.g
	switch qi {
	case 1:
		return gi.x*gj.y > gj.x*gi.y
	case 3:
		return gi.y*gj.x < gj.y*gi.x
	case 5:
		return gi.x*gj.y > gj.x*gi.y
	case 7:
		return gi.y*gj.x < gj.y*gi.x
	default:
		panic(qi)
	}
}

func (v *vectors) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

func P2(pp []point) (nn []point) {
	p, _ := P1(pp)
	m := map[point][]int{}
	for _, q := range pp {
		if p == q {
			continue
		}
		d, g := q.sub(p).div()
		m[g] = append(m[g], d)
	}
	vv := make(vectors, 0, len(m))
	for g, dd := range m {
		sort.Ints(dd)
		vv = append(vv, vector{g, dd})
	}
	sort.Sort(&vv)
	for len(vv) > 0 {
		for i := 0; i < len(vv); {
			v := &vv[i]
			nn, v.dd = append(nn, p.add(v.g.mul(v.dd[0]))), v.dd[1:]
			if len(v.dd) == 0 {
				if i < len(vv)-1 {
					copy(vv[i:], vv[i+1:])
				}
				vv[len(vv)-1] = vector{}
				vv = vv[:len(vv)-1]
			} else {
				i++
			}
		}
	}
	return
}
