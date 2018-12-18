package p

func P1(input []string, n int, stuck bool) int {
	type P struct {
		X, Y int
	}
	adjs := func(p P) []P {
		return []P{
			P{p.X - 1, p.Y - 1},
			P{p.X, p.Y - 1},
			P{p.X + 1, p.Y - 1},
			P{p.X - 1, p.Y},
			P{p.X + 1, p.Y},
			P{p.X - 1, p.Y + 1},
			P{p.X, p.Y + 1},
			P{p.X + 1, p.Y + 1},
		}
	}
	m1, m2 := map[P]struct{}{}, map[P]struct{}{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '#' {
				m1[P{x, y}] = struct{}{}
			}
		}
	}
	for i := 0; i < n; i++ {
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				n, p := 0, P{x, y}
				for _, q := range adjs(p) {
					if _, ok := m1[q]; ok {
						n++
					}
				}
				if _, ok := m1[p]; ok {
					if n == 2 || n == 3 {
						m2[p] = struct{}{}
					} else {
						delete(m2, p)
					}
				} else {
					if n == 3 {
						m2[p] = struct{}{}
					} else {
						delete(m2, p)
					}
				}
			}
		}
		if stuck {
			m2[P{0, 0}] = struct{}{}
			m2[P{0, len(input) - 1}] = struct{}{}
			m2[P{len(input[0]) - 1, 0}] = struct{}{}
			m2[P{len(input[0]) - 1, len(input) - 1}] = struct{}{}
		}
		m1, m2 = m2, m1
	}
	return len(m1)
}
