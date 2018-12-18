package p

import (
	"math"
)

type Position struct {
	X, Y int
}

type Velocity struct {
	X, Y int
}

func (p Position) Add(v Velocity) Position {
	return Position{p.X + v.X, p.Y + v.Y}
}

type PositionVelocity struct {
	Position
	Velocity
}

func P1(input []PositionVelocity, limitY, limitL int) (secs int, output string) {
	m := make(map[Position][]Velocity, len(input))
	for _, pv := range input {
		m[pv.Position] = append(m[pv.Position], pv.Velocity)
	}
	for l := 0; l < limitL; l++ {
		found := false
		for p := range m {
			found = true
			for y := 1; y < limitY; y++ {
				_, ok := m[Position{p.X, p.Y + y}]
				found = found && ok
				if !found {
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
		m2 := make(map[Position][]Velocity, len(m))
		for p, vv := range m {
			for _, v := range vv {
				p2 := p.Add(v)
				m2[p2] = append(m2[p2], v)
			}
		}
		m = m2
		secs++
	}
	minY, maxY := math.MaxInt32, math.MinInt32
	minX, maxX := math.MaxInt32, math.MinInt32
	for p := range m {
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			_, ok := m[Position{x, y}]
			if ok {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return
}
