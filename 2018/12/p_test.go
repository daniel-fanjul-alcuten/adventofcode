package p

import (
	"strings"
	"testing"
)

func TestP1_S(t *testing.T) {
	p := Pots{0, S1}
	if p = p.Gen(S2); p != (Pots{-2, "..#...#....#.....#..#..#..#.."}) {
		t.Error(p)
	}
	if p = p.Gen(S2); p != (Pots{-2, "..##..##...##....#..#..#..##."}) {
		t.Error(p)
	}
	for i := 0; i < 17; i++ {
		p = p.Gen(S2)
	}
	if p = p.Gen(S2); p != (Pots{-4, "..#....##....#####...#######....#.#..##."}) {
		t.Error(p)
	}
	if c := p.Sum(); c != 325 {
		t.Error(c)
	}
}

func TestP1_X(t *testing.T) {
	p := Pots{0, X1}
	for i := 0; i < 20; i++ {
		p = p.Gen(X2)
	}
	if c := p.Sum(); c != 3890 {
		t.Error(c)
	}
}

func TestP2_X(t *testing.T) {
	p := Pots{0, X1}
	for i, l := 0, 50000000000; i < l; i++ {
		q := p.Gen(X2)
		if strings.Trim(p.Line, ".") == strings.Trim(q.Line, ".") {
			for strings.HasPrefix(p.Line, ".") {
				p.Init, p.Line = p.Init+1, p.Line[1:]
			}
			for strings.HasPrefix(q.Line, ".") {
				q.Init, q.Line = q.Init+1, q.Line[1:]
			}
			n := l - (i + 1)
			i += n
			q.Init += n * (q.Init - p.Init)
		}
		p = q
	}
	if c := p.Sum(); c != 4800000001087 {
		t.Error(c)
	}
}
