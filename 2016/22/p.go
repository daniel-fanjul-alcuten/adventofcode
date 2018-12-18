package p

type Node struct {
	X, Y              int
	Size, Used, Avail int
}

func P1(nn []Node) (n int) {
	for i1, n1 := range nn {
		if n1.Used == 0 {
			continue
		}
		for i2, n2 := range nn {
			if i1 == i2 {
				continue
			}
			if n2.Avail < n1.Used {
				continue
			}
			n++
		}
	}
	return
}

func P2(nn []Node) (n int) {
	type P struct {
		X, Y int
	}
	adjs := func(p P) []P {
		return []P{
			{p.X - 1, p.Y},
			{p.X + 1, p.Y},
			{p.X, p.Y - 1},
			{p.X, p.Y + 1},
		}
	}
	type N struct {
		U, A int
	}
	e, d, nodes := P{}, P{}, map[P]N{}
	for _, n := range nn {
		if n.Used == 0 {
			e = P{n.X, n.Y}
		}
		if n.X > d.X {
			d.X = n.X
		}
		nodes[P{n.X, n.Y}] = N{n.Used, n.Avail}
	}
	pathfind := func(source, target P) (path []P) {
		type US struct {
			P P
			A int
		}
		type S struct {
			P P
			A int
			X []P
		}
		s := S{source, nodes[source].A, nil}
		states, processed := []S{s}, map[US]struct{}{{s.P, s.A}: struct{}{}}
		for len(states) > 0 {
			s, states = states[0], states[1:]
			if s.P == target {
				if len(path) == 0 || len(s.X) < len(path) {
					path = append(path[:0], s.X...)
				}
				continue
			}
			if len(path) != 0 && len(s.X) >= len(path)-2 {
				continue
			}
			for _, p := range adjs(s.P) {
				if p == d {
					continue
				}
				np, ok := nodes[p]
				if !ok {
					continue
				}
				if np.U > s.A {
					continue
				}
				a := s.A
				if na := np.U + np.A; na < a {
					a = na
				}
				if _, ok := processed[US{p, a}]; ok {
					continue
				}
				states, processed[US{p, a}] = append(states, S{p, a,
					append(append(make([]P, 0, len(s.X)), s.X...), p)}), struct{}{}
			}
		}
		return
	}
	move := func(source, target P) {
		ns, ok := nodes[source]
		if !ok {
			panic(source)
		}
		nt, ok := nodes[target]
		if !ok {
			panic(target)
		}
		if ns.U > nt.A {
			panic(ns.U)
		}
		nt.A -= ns.U
		nt.U += ns.U
		ns.A += ns.U
		ns.U -= ns.U
		nodes[source] = ns
		nodes[target] = nt
		if e == target {
			e = source
		}
		if d == source {
			d = target
		}
		n++
	}
	for {
		if d == (P{}) {
			break
		}
		if e != (P{d.X - 1, d.Y}) {
			x := pathfind(e, P{d.X - 1, d.Y})
			if len(x) == 0 {
				panic(x)
			}
			for _, p := range x {
				move(p, e)
			}
			move(d, e)
		}
	}
	return
}
