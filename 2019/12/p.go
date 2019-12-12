package p

import (
	"fmt"
)

type pos struct {
	x, y, z int
}

type vel struct {
	x, y, z int
}

func (p pos) add(v vel) pos {
	return pos{p.x + v.x, p.y + v.y, p.z + v.z}
}

type moon struct {
	p pos
	v vel
}

func P1(time int, pp ...pos) (e int) {
	moons := make([]*moon, len(pp))
	for i, p := range pp {
		moons[i] = &moon{pos{p.x, p.y, p.z}, vel{}}
	}
	for s := 0; s < time; s++ {
		for i, mi := range moons {
			for j, mj := range moons {
				if i == j {
					continue
				}
				if mi.p.x > mj.p.x {
					mi.v.x--
				} else if mi.p.x < mj.p.x {
					mi.v.x++
				}
				if mi.p.y > mj.p.y {
					mi.v.y--
				} else if mi.p.y < mj.p.y {
					mi.v.y++
				}
				if mi.p.z > mj.p.z {
					mi.v.z--
				} else if mi.p.z < mj.p.z {
					mi.v.z++
				}
			}
		}
		for _, m := range moons {
			m.p = m.p.add(m.v)
		}
	}
	for _, m := range moons {
		p := 0
		if m.p.x > 0 {
			p += m.p.x
		} else {
			p -= m.p.x
		}
		if m.p.y > 0 {
			p += m.p.y
		} else {
			p -= m.p.y
		}
		if m.p.z > 0 {
			p += m.p.z
		} else {
			p -= m.p.z
		}
		k := 0
		if m.v.x > 0 {
			k += m.v.x
		} else {
			k -= m.v.x
		}
		if m.v.y > 0 {
			k += m.v.y
		} else {
			k -= m.v.y
		}
		if m.v.z > 0 {
			k += m.v.z
		} else {
			k -= m.v.z
		}
		e += p * k
	}
	return
}

type moon2 struct {
	p, v int
}

func p2(pp ...int) (t int) {
	moons := make([]*moon2, len(pp))
	for i, p := range pp {
		moons[i] = &moon2{p, 0}
	}
	state := func() (s string) {
		for _, m := range moons {
			s += fmt.Sprintf("%v:%v,", m.p, m.v)
		}
		return
	}
	states := map[string]struct{}{state(): struct{}{}}
	for t++; ; t++ {
		for i, mi := range moons {
			for j, mj := range moons {
				if i == j {
					continue
				}
				if mi.p > mj.p {
					mi.v--
				} else if mi.p < mj.p {
					mi.v++
				}
			}
		}
		for _, m := range moons {
			m.p += m.v
		}
		s := state()
		if _, ok := states[s]; ok {
			return
		}
		states[s] = struct{}{}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func P2(pp ...pos) (t int) {
	ii := make([]int, len(pp))
	for i, p := range pp {
		ii[i] = p.x
	}
	tx := p2(ii...)
	for i, p := range pp {
		ii[i] = p.y
	}
	ty := p2(ii...)
	t = lcm(tx, ty)
	for i, p := range pp {
		ii[i] = p.z
	}
	tz := p2(ii...)
	t = lcm(t, tz)
	return
}
