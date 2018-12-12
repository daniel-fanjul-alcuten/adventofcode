package p

import "strings"

type Pots struct {
	Init int
	Line string
}

func (p Pots) Gen(rules map[string]string) (q Pots) {
	for !strings.HasPrefix(p.Line, "....") {
		p.Init, p.Line = p.Init-1, "."+p.Line
	}
	for !strings.HasSuffix(p.Line, "....") {
		p.Line = p.Line + "."
	}
	min, max := 2, len(p.Line)-2
	q.Init = p.Init + 2
	for i := min; i < max; i++ {
		in := p.Line[i-2 : i+3]
		out := rules[in]
		if out == "" {
			out = "."
		}
		q.Line += out
	}
	return
}

func (p Pots) Sum() (s int) {
	for i := 0; i < len(p.Line); i++ {
		if p.Line[i:i+1] == "#" {
			s += p.Init + i
		}
	}
	return
}
