package p

import (
	"math"
	"strconv"
	"strings"
)

type P struct {
	X, Y int
}

func P12(input []string) (p1, p2 int) {
	clay := map[P]bool{}
	for _, in := range input {
		i := strings.Index(in, ",")
		x, err := strconv.Atoi(in[2:i])
		if err != nil {
			panic(err)
		}
		j := i + 4 + strings.Index(in[i+4:], ".")
		y0, err := strconv.Atoi(in[i+4 : j])
		if err != nil {
			panic(err)
		}
		y1, err := strconv.Atoi(in[j+2:])
		if err != nil {
			panic(err)
		}
		if in[0] == 'x' {
			for y := y0; y <= y1; y++ {
				clay[P{x, y}] = true
			}
			continue
		}
		y, x0, x1 := x, y0, y1
		for x := x0; x <= x1; x++ {
			clay[P{x, y}] = true
		}
	}
	minY, maxY := math.MaxInt32, math.MinInt32
	for p := range clay {
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	fountain := P{500, 0}
	fountains := []P{fountain}
	watered := map[P]bool{}
	for len(fountains) > 0 {
		last := len(fountains) - 1
		fountain, fountains = fountains[last], fountains[:last]
		if clay[fountain] {
			continue
		}
		p := fountain
		for p.Y+1 <= maxY && !clay[P{p.X, p.Y + 1}] {
			p.Y++
			watered[p] = true
		}
		if !clay[P{p.X, p.Y + 1}] {
			continue
		}
		enclosed := true
		l := p
		for !clay[P{l.X - 1, l.Y}] && clay[P{l.X - 1, l.Y + 1}] {
			l.X--
			watered[l] = true
		}
		if !clay[P{l.X - 1, l.Y}] {
			l.X--
			enclosed = false
			if !watered[l] {
				watered[l] = true
				fountains = append(fountains, fountain, l)
			}
		}
		r := p
		for !clay[P{r.X + 1, r.Y}] && clay[P{r.X + 1, r.Y + 1}] {
			r.X++
			watered[r] = true
		}
		if !clay[P{r.X + 1, r.Y}] {
			r.X++
			if enclosed {
				enclosed = false
				if !watered[r] {
					watered[r] = true
					fountains = append(fountains, fountain, r)
				}
			} else {
				if !watered[r] {
					watered[r] = true
					fountains = append(fountains, r)
				}
			}
		}
		if enclosed {
			for x := l.X; x <= r.X; x++ {
				clay[P{x, p.Y}] = true
			}
			fountains = append(fountains, fountain)
		}
	}
	minX, maxX := math.MaxInt32, math.MinInt32
	for p := range watered {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := P{x, y}
			if watered[p] {
				p1++
				if clay[p] {
					p2++
				}
			}
		}
	}
	return
}
