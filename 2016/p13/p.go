package p

import (
	"math/bits"
)

type Position struct {
	X, Y int
}

func (p Position) Adjacents() []Position {
	return []Position{
		{p.X, p.Y - 1},
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
		{p.X, p.Y + 1},
	}
}

type State struct {
	P Position
	N int
}

func P1(input int, target Position) int {
	open := func(x, y int) bool {
		return bits.OnesCount(uint(x*x+3*x+2*x*y+y+y*y+input))%2 == 0
	}
	pending, processed := []State{}, map[Position]struct{}{}
	s := State{Position{1, 1}, 0}
	pending, processed[s.P] = append(pending, s), struct{}{}
	for len(pending) > 0 {
		s, pending = pending[0], pending[1:]
		for _, p := range s.P.Adjacents() {
			if p.X < 0 {
				continue
			}
			if p.Y < 0 {
				continue
			}
			if !open(p.X, p.Y) {
				continue
			}
			if p == target {
				return s.N + 1
			}
			if _, ok := processed[p]; ok {
				continue
			}
			t := State{p, s.N + 1}
			pending, processed[t.P] = append(pending, t), struct{}{}
		}
	}
	return -1
}

func P2(input int, steps int) int {
	open := func(x, y int) bool {
		return bits.OnesCount(uint(x*x+3*x+2*x*y+y+y*y+input))%2 == 0
	}
	pending, processed := []State{}, map[Position]struct{}{}
	s := State{Position{1, 1}, 0}
	pending, processed[s.P] = append(pending, s), struct{}{}
	for len(pending) > 0 {
		s, pending = pending[0], pending[1:]
		if s.N >= steps {
			continue
		}
		for _, p := range s.P.Adjacents() {
			if p.X < 0 {
				continue
			}
			if p.Y < 0 {
				continue
			}
			if !open(p.X, p.Y) {
				continue
			}
			if _, ok := processed[p]; ok {
				continue
			}
			t := State{p, s.N + 1}
			pending, processed[t.P] = append(pending, t), struct{}{}
		}
	}
	return len(processed)
}
