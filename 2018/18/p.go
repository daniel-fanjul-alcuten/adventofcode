package p

import (
	"strings"
)

func P1(input []string, rounds int) (output1, output2 int) {
	type P struct {
		X, Y int
	}
	adjs := func(p P) []P {
		return []P{
			{p.X - 1, p.Y - 1},
			{p.X, p.Y - 1},
			{p.X + 1, p.Y - 1},
			{p.X - 1, p.Y},
			{p.X + 1, p.Y},
			{p.X - 1, p.Y + 1},
			{p.X, p.Y + 1},
			{p.X + 1, p.Y + 1},
		}
	}
	m1, m2 := map[P]byte{}, map[P]byte{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			m1[P{x, y}] = input[y][x]
		}
	}
	state := func() string {
		b := strings.Builder{}
		for y := 0; y < len(input); y++ {
			for x := 0; x < len(input[y]); x++ {
				b.WriteRune(rune(m1[P{x, y}]))
			}
			b.WriteRune('\n')
		}
		return b.String()
	}
	states := map[string]int{}
	states[state()] = 0
	cycle := 0
	for r := 1; r <= rounds; r++ {
		for p := range m1 {
			open, wood, lumber := 0, 0, 0
			for _, q := range adjs(p) {
				switch b := m1[q]; b {
				case 0, '.':
					open++
				case '|':
					wood++
				case '#':
					lumber++
				default:
					panic(b)
				}
			}
			switch b := m1[p]; b {
			case '.':
				if wood >= 3 {
					m2[p] = '|'
					break
				}
				m2[p] = '.'
			case '|':
				if lumber >= 3 {
					m2[p] = '#'
					break
				}
				m2[p] = '|'
			case '#':
				if wood >= 1 && lumber >= 1 {
					m2[p] = '#'
					break
				}
				m2[p] = '.'
			default:
				panic(b)
			}
		}
		for p := range m1 {
			delete(m1, p)
		}
		m1, m2 = m2, m1
		if cycle == 0 {
			s := state()
			r2, ok := states[s]
			if !ok {
				states[s] = r
				continue
			}
			cycle = r - r2
			r += cycle * ((rounds - r) / cycle)
		}
	}
	wood, lumber := 0, 0
	for p := range m1 {
		switch b := m1[p]; b {
		case '.':
		case '|':
			wood++
		case '#':
			lumber++
		default:
			panic(b)
		}
	}
	output1, output2 = wood, lumber
	return
}
