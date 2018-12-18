package p1

import (
	"strconv"
	"strings"
)

type P struct {
	X, Y int
}

type Pos struct {
	P, V P
}

func New() Pos {
	return Pos{
		P{0, 0},
		P{0, 1},
	}
}

func (p *Pos) Right() {
	switch p.V {
	case P{0, 1}:
		p.V = P{1, 0}
	case P{1, 0}:
		p.V = P{0, -1}
	case P{0, -1}:
		p.V = P{-1, 0}
	case P{-1, 0}:
		p.V = P{0, 1}
	}
}

func (p *Pos) Left() {
	switch p.V {
	case P{0, 1}:
		p.V = P{-1, 0}
	case P{-1, 0}:
		p.V = P{0, -1}
	case P{0, -1}:
		p.V = P{1, 0}
	case P{1, 0}:
		p.V = P{0, 1}
	}
}

func (p *Pos) Move(n int) {
	p.P.X += n * p.V.X
	p.P.Y += n * p.V.Y
}

func (p *Pos) Exec1(s string) {
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(err)
	}
	if s[0] == 'R' {
		p.Right()
		p.Move(n)
	} else if s[0] == 'L' {
		p.Left()
		p.Move(n)
	} else {
		panic(s[0])
	}
}

func (p *Pos) Exec1N(ss string) {
	for _, s := range strings.Split(ss, ", ") {
		p.Exec1(s)
	}
}

func (p *Pos) Distance() (d int) {
	if p.P.X < 0 {
		d -= p.P.X
	} else {
		d += p.P.X
	}
	if p.P.Y < 0 {
		d -= p.P.Y
	} else {
		d += p.P.Y
	}
	return
}

func (p *Pos) Move2(n int, m map[P]struct{}) bool {
	for i := 0; i < n; i++ {
		p.P.X += p.V.X
		p.P.Y += p.V.Y
		if _, ok := m[p.P]; ok {
			return true
		}
		m[p.P] = struct{}{}
	}
	return false
}

func (p *Pos) Exec2(s string, m map[P]struct{}) bool {
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(err)
	}
	if s[0] == 'R' {
		p.Right()
		return p.Move2(n, m)
	} else if s[0] == 'L' {
		p.Left()
		return p.Move2(n, m)
	} else {
		panic(s[0])
	}
}

func (p *Pos) Exec2N(ss string, m map[P]struct{}) bool {
	for _, s := range strings.Split(ss, ", ") {
		if p.Exec2(s, m) {
			return true
		}
	}
	return false
}
