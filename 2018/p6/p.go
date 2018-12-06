package p

import (
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Distance(qq []Point) (distance int, points []Point) {
	distance = math.MaxInt64
	for _, q := range qq {
		d := 0
		if p.X >= q.X {
			d += p.X - q.X
		} else {
			d += q.X - p.X
		}
		if p.Y >= q.Y {
			d += p.Y - q.Y
		} else {
			d += q.Y - p.Y
		}
		if d < distance {
			distance, points = d, append(points[:0], q)
		} else if d == distance {
			points = append(points, q)
		}
	}
	return
}

func P1(pp []Point) (maxArea int) {
	maxX, maxY := 0, 0
	for _, p := range pp {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	maxX, maxY = 2*maxX, 2*maxY
	p, m, e := Point{}, map[Point]int{}, map[Point]struct{}{}
	for p.X = -maxX; p.X <= maxX; p.X++ {
		for p.Y = -maxY; p.Y <= maxY; p.Y++ {
			_, qq := p.Distance(pp)
			if len(qq) == 1 {
				m[qq[0]]++
				if p.X == -maxX || p.X == maxX || p.Y == -maxY || p.Y == maxY {
					for _, q := range qq {
						e[q] = struct{}{}
					}
				}
			}
		}
	}
	for q := range e {
		delete(m, q)
	}
	maxArea = 0
	for _, a := range m {
		if a > maxArea {
			maxArea = a
		}
	}
	return
}

func P2(maxD int, pp []Point) (maxArea int) {
	maxX, maxY := 0, 0
	for _, p := range pp {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	maxX, maxY = 2*maxX, 2*maxY
	p := Point{}
	for p.X = -maxX; p.X <= maxX; p.X++ {
		for p.Y = -maxY; p.Y <= maxY; p.Y++ {
			pd := 0
			for i := range pp {
				d, _ := p.Distance(pp[i : i+1])
				pd += d
			}
			if pd < maxD {
				maxArea++
			}
		}
	}
	return
}
