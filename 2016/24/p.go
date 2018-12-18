package p

type P struct {
	X, Y int
}

func (p P) Adjs() []P {
	return []P{
		P{p.X + 1, p.Y},
		P{p.X - 1, p.Y},
		P{p.X, p.Y + 1},
		P{p.X, p.Y - 1},
	}
}

func P1(input []string, retTo0 bool) (output int) {
	b2p := map[byte]P{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			b := input[y][x]
			if b >= '0' && b <= '9' {
				b2p[b] = P{x, y}
			}
		}
	}
	pathfind := func(source, target P) (n int) {
		type S struct {
			P P
			L int
		}
		s := S{source, 0}
		states, processed := []S{s}, map[P]struct{}{s.P: struct{}{}}
		for len(states) > 0 {
			s, states = states[0], states[1:]
			if s.P == target {
				if n == 0 || s.L < n {
					n = s.L
				}
				continue
			}
			if n != 0 && s.L >= n-2 {
				continue
			}
			for _, p := range s.P.Adjs() {
				if input[p.Y][p.X] == '#' {
					continue
				}
				if _, ok := processed[p]; ok {
					continue
				}
				states, processed[p] = append(states, S{p, s.L + 1}), struct{}{}
			}
		}
		return
	}
	type Path struct {
		S, T byte
	}
	paths := map[Path]int{}
	for bs, ps := range b2p {
		for bt, pt := range b2p {
			if bs == bt {
				continue
			}
			if paths[Path{bs, bt}] != 0 {
				continue
			}
			n := pathfind(ps, pt)
			if n == 0 {
				panic(n)
			}
			paths[Path{bs, bt}] = n
			paths[Path{bt, bs}] = n
		}
	}
	var rec func(b byte, m map[byte]struct{}) int
	rec = func(b byte, m map[byte]struct{}) int {
		if retTo0 {
			if len(m) == 0 {
				return paths[Path{b, '0'}]
			}
		} else {
			if len(m) == 1 {
				for c := range m {
					return paths[Path{b, c}]
				}
			}
		}
		cc := make([]byte, 0, len(m))
		for c := range m {
			cc = append(cc, c)
		}
		bestN := 0
		for _, c := range cc {
			delete(m, c)
			n := paths[Path{b, c}] + rec(c, m)
			m[c] = struct{}{}
			if bestN == 0 || n < bestN {
				bestN = n
			}
		}
		return bestN
	}
	if !retTo0 {
	}
	m := map[byte]struct{}{}
	for b := range b2p {
		if b != '0' {
			m[b] = struct{}{}
		}
	}
	output = rec('0', m)
	return
}
