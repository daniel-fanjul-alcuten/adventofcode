package p11

import (
	"strings"
)

type Position struct {
	X, Y int
	D    int
}

func (p Position) Add(q Position) (r Position) {
	r = Position{p.X + q.X, p.Y + q.Y, p.D}
	if d := r.Distance(); d > p.D {
		r.D = d
	}
	return
}

const (
	N  = "n"
	S  = "s"
	NE = "ne"
	SE = "se"
	NW = "nw"
	SW = "sw"
)

var M = map[string]Position{
	N:  {0, 2, 0},
	S:  {0, -2, 0},
	NE: {1, 1, 0},
	SE: {1, -1, 0},
	NW: {-1, 1, 0},
	SW: {-1, -1, 0},
}

func (p Position) Add1(s string) (r Position) {
	return p.Add(M[s])
}

func (p Position) AddN(ss string) (r Position) {
	r = p
	for _, s := range strings.Split(ss, ",") {
		r = r.Add1(s)
	}
	return
}

func (p Position) Distance() (r int) {
	x, y := p.X, p.Y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	y = y/2 + y%2
	if x > y {
		r = x
	} else {
		r = y
	}
	return
}
