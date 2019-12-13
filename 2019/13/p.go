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

type pos struct {
	x, y int
}

func P1(code []int) (n int) {
	out := New(code).Run()
	m := map[pos]int{}
	for i := 0; i < len(out); i += 3 {
		x := out[i+0]
		y := out[i+1]
		v := out[i+2]
		m[pos{x, y}] = v
	}
	for _, v := range m {
		if v == 2 {
			n++
		}
	}
	return
}

func P2(code []int) int {
	i := New(code)
	i.s[0] = 2
	score, blocks, paddle, ball, m := 0, 0, pos{}, pos{}, map[pos]int{}
	run := func(input ...int) {
		out := i.Run(input...)
		for i := 0; i < len(out); i += 3 {
			x := out[i+0]
			y := out[i+1]
			v := out[i+2]
			if x == -1 && y == 0 {
				score = v
			}
			if m[pos{x, y}] == 2 {
				blocks--
			}
			if v == 2 {
				blocks++
			} else if v == 3 {
				paddle = pos{x, y}
			} else if v == 4 {
				ball = pos{x, y}
			}
			m[pos{x, y}] = v
		}
	}
	run()
	for blocks > 0 {
		switch {
		case ball.x < paddle.x:
			run(-1)
		case ball.x > paddle.x:
			run(+1)
		case ball.x == paddle.x:
			run(0)
		}
	}
	return score
}
