package p

import (
	"fmt"
	"sort"
)

type Floor []string

func (f Floor) Safe() bool {
	gen := false
	for _, g := range f {
		if gen = g[1] == 'G'; gen {
			break
		}
	}
	if !gen {
		return true
	}
	mic := false
	for _, m := range f {
		if mic = m[1] == 'M'; mic {
			break
		}
	}
	if !mic {
		return true
	}
	for _, m := range f {
		if m[1] == 'M' {
			ok := false
			for _, g := range f {
				if g[1] == 'G' {
					if ok = g[0] == m[0]; ok {
						break
					}
				}
			}
			if !ok {
				return false
			}
		}
	}
	return true
}

type Floors []Floor

func (ff Floors) String() string {
	gg := make([]Floor, 0, len(ff))
	for _, f := range ff {
		g := append(make(Floor, 0, len(f)), f...)
		sort.Strings(g)
		gg = append(gg, g)
	}
	return fmt.Sprint(gg)
}

func (ff Floors) Produce(sf, tf int) (gg []Floors) {
	var sf2, tf2 Floor
	for i := range ff[sf] {
		sf2 := append(append(sf2[:0],
			ff[sf][:i]...),
			ff[sf][i+1:]...)
		if sf2.Safe() {
			tf2 := append(append(tf2[:0],
				ff[tf]...),
				ff[sf][i])
			if tf2.Safe() {
				g := append(make(Floors, 0, len(ff)), ff...)
				g[sf], sf2 = sf2, nil
				g[tf], tf2 = tf2, nil
				gg = append(gg, g)
			}
		}
	}
	for i := range ff[sf] {
		for j := i + 1; j < len(ff[sf]); j++ {
			sf2 := append(append(append(sf2[:0],
				ff[sf][:i]...),
				ff[sf][i+1:j]...),
				ff[sf][j+1:]...)
			if sf2.Safe() {
				tf2 := append(append(tf2[:0],
					ff[tf]...),
					ff[sf][i],
					ff[sf][j],
				)
				if tf2.Safe() {
					g := append(make(Floors, 0, len(ff)), ff...)
					g[sf], sf2 = sf2, nil
					g[tf], tf2 = tf2, nil
					gg = append(gg, g)
				}
			}
		}
	}
	return
}

type Puzzle struct {
	P, N int
	F    Floors
}

func (p Puzzle) CanonicalString() string {
	g, m := make(Floors, len(p.F)), map[string]int{}
	for i, f := range p.F {
		for _, t := range f {
			n := m[t[:1]]
			if n == 0 {
				n = len(m) + 1
				m[t[:1]] = n
			}
			g[i] = append(g[i], fmt.Sprint(n)+t[1:])
		}
	}
	return fmt.Sprintf("%v@%v", p.P, g)
}

func (p Puzzle) String() string {
	return fmt.Sprintf("%v@%v+%v", p.P, p.F, p.N)
}

func (p Puzzle) Solved() bool {
	for i := 0; i < len(p.F)-1; i++ {
		if len(p.F[i]) > 0 {
			return false
		}
	}
	return p.P == len(p.F)-1
}

func (p Puzzle) Produce() (qq []Puzzle) {
	if p.P > 0 {
		for _, f := range p.F.Produce(p.P, p.P-1) {
			qq = append(qq, Puzzle{p.P - 1, p.N + 1, f})
		}
	}
	if p.P < len(p.F)-1 {
		for _, f := range p.F.Produce(p.P, p.P+1) {
			qq = append(qq, Puzzle{p.P + 1, p.N + 1, f})
		}
	}
	return
}

func (p Puzzle) P2() (n int) {
	pending, processed := []Puzzle{}, map[string]struct{}{}
	if p.Solved() {
		return p.N
	}
	n, pending, processed[p.CanonicalString()] = -1, append(pending, p), struct{}{}
	for len(pending) > 0 {
		p, pending = pending[0], pending[1:]
		for _, q := range p.Produce() {
			if _, ok := processed[q.CanonicalString()]; ok {
				continue
			}
			if q.Solved() {
				if n == -1 || q.N < n {
					n = q.N
				}
				continue
			}
			if n == -1 || q.N < n-1 {
				pending = append(pending, q)
				processed[q.CanonicalString()] = struct{}{}
			}
		}
	}
	return
}
