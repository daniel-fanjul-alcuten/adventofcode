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

func (p pos) moves() [4]pos {
	return [4]pos{
		pos{p.x, p.y - 1},
		pos{p.x, p.y + 1},
		pos{p.x - 1, p.y},
		pos{p.x + 1, p.y},
	}
}

type node struct {
	p      pos
	ic     *Intcode
	inputs int
	output int
}

func P1(code []int) int {
	visited := map[pos]struct{}{}
	pending := []node{node{pos{}, New(code), 0, -1}}
	for len(pending) > 0 {
		n := pending[0]
		pending = pending[1:]
		if _, ok := visited[n.p]; ok {
			continue
		}
		visited[n.p] = struct{}{}
		for i, p := range n.p.moves() {
			if _, ok := visited[p]; ok {
				continue
			}
			ic := n.ic.clone()
			inputs := n.inputs + 1
			outputs := ic.Run(i + 1)
			if len(outputs) != 1 {
				panic(len(outputs))
			}
			output := outputs[len(outputs)-1]
			if output == 0 {
				visited[p] = struct{}{}
				continue
			} else if output == 2 {
				return inputs
			}
			pending = append(pending, node{p, ic, inputs, output})
		}
	}
	return -1
}

func P2(code []int) int {
	visited := map[pos]int{}
	pending := []node{node{pos{}, New(code), 0, -1}}
	for len(pending) > 0 {
		n := pending[0]
		pending = pending[1:]
		if _, ok := visited[n.p]; ok {
			continue
		}
		visited[n.p] = n.output
		for i, p := range n.p.moves() {
			if _, ok := visited[p]; ok {
				continue
			}
			ic := n.ic.clone()
			inputs := n.inputs + 1
			outputs := ic.Run(i + 1)
			if len(outputs) != 1 {
				panic(len(outputs))
			}
			output := outputs[len(outputs)-1]
			if output == 0 {
				visited[p] = output
				continue
			}
			pending = append(pending, node{p, ic, inputs, output})
		}
	}
	filled, filled2, empty := []pos{}, []pos{}, map[pos]struct{}{}
	for p, o := range visited {
		switch o {
		case 1:
			empty[p] = struct{}{}
		case 2:
			filled = append(filled, p)
		}
	}
	t := 0
	for len(empty) > 0 {
		for _, p := range filled {
			for _, q := range p.moves() {
				if _, ok := empty[q]; ok {
					delete(empty, q)
					filled2 = append(filled2, q)
				}
			}
		}
		filled, filled2 = filled2, filled[:0]
		t++
	}
	return t
}
