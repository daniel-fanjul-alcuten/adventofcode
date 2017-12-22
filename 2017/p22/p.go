package p22

type P struct {
	X, Y int
}

func (p P) Add(q P) P {
	return P{p.X + q.X, p.Y + q.Y}
}

func (p P) Right() P {
	switch p {
	case P{0, -1}:
		return P{1, 0}
	case P{1, 0}:
		return P{0, 1}
	case P{0, 1}:
		return P{-1, 0}
	case P{-1, 0}:
		return P{0, -1}
	default:
		panic(p)
	}
}

func (p P) Left() (q P) {
	switch p {
	case P{0, -1}:
		return P{-1, 0}
	case P{-1, 0}:
		return P{0, 1}
	case P{0, 1}:
		return P{1, 0}
	case P{1, 0}:
		return P{0, -1}
	default:
		panic(p)
	}
}

type Grid struct {
	Inf    map[P]byte
	Pos    P
	Dir    P
	NumBur int
	NumInf int
}

func (g *Grid) Parse(ss []string) {
	g.Inf = map[P]byte{}
	for y, s := range ss {
		for x, b := range []byte(s) {
			if b == '#' {
				g.Inf[P{x, y}] = b
			}
		}
	}
	c := len(ss) / 2
	g.Pos = P{c, c}
	g.Dir = P{0, -1}
	return
}

func (g *Grid) Burst() {
	g.NumBur += 1
	_, f := g.Inf[g.Pos]
	if f {
		g.Dir = g.Dir.Right()
	} else {
		g.Dir = g.Dir.Left()
	}
	if f {
		delete(g.Inf, g.Pos)
	} else {
		g.NumInf += 1
		g.Inf[g.Pos] = '#'
	}
	g.Pos = g.Pos.Add(g.Dir)
}

func (g *Grid) Burst2() {
	g.NumBur += 1
	f, ok := g.Inf[g.Pos]
	if !ok {
		g.Dir = g.Dir.Left()
	} else if f == 'W' {
	} else if f == '#' {
		g.Dir = g.Dir.Right()
	} else if f == 'F' {
		g.Dir = g.Dir.Right().Right()
	} else {
		panic(string(f))
	}
	if !ok {
		g.Inf[g.Pos] = 'W'
	} else if f == 'W' {
		g.NumInf += 1
		g.Inf[g.Pos] = '#'
	} else if f == '#' {
		g.Inf[g.Pos] = 'F'
	} else if f == 'F' {
		delete(g.Inf, g.Pos)
	} else {
		panic(f)
	}
	g.Pos = g.Pos.Add(g.Dir)
}
