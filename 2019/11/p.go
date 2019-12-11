package p

type Intcode struct {
	p, r int
	s    map[int]int
}

func New(code []int) *Intcode {
	p, r, s := 0, 0, make(map[int]int, len(code))
	for i, c := range code {
		s[i] = c
	}
	return &Intcode{p, r, s}
}

func (i *Intcode) v1() int {
	switch i.s[i.p] / 100 % 10 {
	case 0:
		return i.s[i.s[i.p+1]]
	case 1:
		return i.s[i.p+1]
	case 2:
		return i.s[i.r+i.s[i.p+1]]
	default:
		panic(i.s[i.p] / 100 % 10)
	}
}

func (i *Intcode) v2() int {
	switch i.s[i.p] / 1000 % 10 {
	case 0:
		return i.s[i.s[i.p+2]]
	case 1:
		return i.s[i.p+2]
	case 2:
		return i.s[i.r+i.s[i.p+2]]
	default:
		panic(i.s[i.p] / 1000 % 10)
	}
}

func (i *Intcode) a1(v int) {
	switch i.s[i.p] / 100 % 10 {
	case 0:
		i.s[i.s[i.p+1]] = v
	case 2:
		i.s[i.r+i.s[i.p+1]] = v
	default:
		panic(i.s[i.p] / 100 % 10)
	}
}

func (i *Intcode) a3(v int) {
	switch i.s[i.p] / 10000 % 10 {
	case 0:
		i.s[i.s[i.p+3]] = v
	case 2:
		i.s[i.r+i.s[i.p+3]] = v
	default:
		panic(i.s[i.p] / 10000 % 10)
	}
}

func (i *Intcode) Run(input ...int) (output []int) {
	for {
		switch i.s[i.p] % 100 {
		case 1:
			i.a3(i.v1() + i.v2())
			i.p += 4
		case 2:
			i.a3(i.v1() * i.v2())
			i.p += 4
		case 3:
			if len(input) == 0 {
				return
			}
			i.a1(input[0])
			input = input[1:]
			i.p += 2
		case 4:
			output = append(output, i.v1())
			i.p += 2
		case 5:
			if i.v1() != 0 {
				i.p = i.v2()
				break
			}
			i.p += 3
		case 6:
			if i.v1() == 0 {
				i.p = i.v2()
				break
			}
			i.p += 3
		case 7:
			if i.v1() < i.v2() {
				i.a3(1)
			} else {
				i.a3(0)
			}
			i.p += 4
		case 8:
			if i.v1() == i.v2() {
				i.a3(1)
			} else {
				i.a3(0)
			}
			i.p += 4
		case 9:
			i.r += i.v1()
			i.p += 2
		case 99:
			return
		default:
			panic(i.s[i.p] % 100)
		}
	}
}

type point struct {
	x, y int
}

func (p point) right() point {
	switch p {
	case point{0, -1}:
		return point{1, 0}
	case point{1, 0}:
		return point{0, 1}
	case point{0, 1}:
		return point{-1, 0}
	case point{-1, 0}:
		return point{0, -1}
	default:
		panic(p)
	}
}

func (p point) left() point {
	switch p {
	case point{0, -1}:
		return point{-1, 0}
	case point{-1, 0}:
		return point{0, 1}
	case point{0, 1}:
		return point{1, 0}
	case point{1, 0}:
		return point{0, -1}
	default:
		panic(p)
	}
}

func (p point) add(q point) point {
	return point{p.x + q.x, p.y + q.y}
}

func P1(code []int) int {
	p, d, c, x := point{}, point{0, -1}, map[point]int{}, map[point]struct{}{}
	i := New(code)
	for {
		output := i.Run(c[p])
		if len(output) == 0 {
			break
		}
		if len(output) != 2 {
			panic(len(output))
		}
		bc, bd := output[0], output[1]
		c[p], x[p] = bc, struct{}{}
		switch bd {
		case 0:
			d = d.left()
		case 1:
			d = d.right()
		default:
			panic(bd)
		}
		p = p.add(d)
	}
	return len(x)
}

func P2(code []int) (output []string) {
	p, d, c := point{}, point{0, -1}, map[point]int{point{}: 1}
	i := New(code)
	for {
		out := i.Run(c[p])
		if len(out) == 0 {
			break
		}
		if len(out) != 2 {
			panic(len(out))
		}
		c[p] = out[0]
		switch out[1] {
		case 0:
			d = d.left()
		case 1:
			d = d.right()
		default:
			panic(out[1])
		}
		p = p.add(d)
	}
	min, max := point{}, point{}
	for p := range c {
		if p.x < min.x {
			min.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		}
		if p.x > max.x {
			max.x = p.x
		}
		if p.y > max.y {
			max.y = p.y
		}
	}
	for y := min.y; y <= max.y; y++ {
		s := ""
		for x := min.x; x <= max.x; x++ {
			switch c := c[point{x, y}]; c {
			case 0:
				s += " "
			case 1:
				s += "#"
			default:
				panic(c)
			}
		}
		output = append(output, s)
	}
	return
}
