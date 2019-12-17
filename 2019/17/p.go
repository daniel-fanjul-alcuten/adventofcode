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

func (i *Intcode) clone() *Intcode {
	s := make(map[int]int, len(i.s))
	for i, c := range i.s {
		s[i] = c
	}
	return &Intcode{i.p, i.r, s}
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

type pos struct {
	x, y int
}

func tomap(ss []string) (m map[pos]rune) {
	m = map[pos]rune{}
	for y, s := range ss {
		for x, r := range s {
			m[pos{x, y}] = r
		}
	}
	return
}

func parse(code []int) (m map[pos]rune) {
	m = map[pos]rune{}
	p, out := pos{}, New(code).Run()
	for _, o := range out {
		if o == 10 {
			p = pos{0, p.y + 1}
		} else {
			m[p] = rune(o)
			p.x++
		}
	}
	return
}

func P1(m map[pos]rune) (v int) {
	for p := range m {
		if m[pos{p.x, p.y}] == '#' &&
			m[pos{p.x - 1, p.y}] == '#' &&
			m[pos{p.x + 1, p.y}] == '#' &&
			m[pos{p.x, p.y - 1}] == '#' &&
			m[pos{p.x, p.y + 1}] == '#' {
			v += p.x * p.y
		}
	}
	return
}
