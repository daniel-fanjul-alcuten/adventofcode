package p12

import (
	"strings"
)

type Program struct {
	Source  string
	Targets map[string]struct{}
}

func Parse1(s string) (p Program) {
	ss := strings.Split(s, " <-> ")
	if l := len(ss); l != 2 {
		panic(l)
	}
	p.Source = ss[0]
	tt := strings.Split(ss[1], ", ")
	if l := len(tt); l < 1 {
		panic(l)
	}
	p.Targets = make(map[string]struct{}, len(tt))
	for _, t := range tt {
		p.Targets[t] = struct{}{}
	}
	return
}

type Programs map[string]Program

func ParseN(ss []string) (m Programs) {
	m = make(Programs, len(ss))
	for _, s := range ss {
		p := Parse1(s)
		m[p.Source] = p
	}
	return
}

func (pp Programs) Group(s string) (g map[string]struct{}) {
	pending := map[string]struct{}{s: struct{}{}}
	g = map[string]struct{}{}
	for len(pending) > 0 {
		for p := range pending {
			if _, ok := g[p]; !ok {
				for t := range pp[p].Targets {
					pending[t] = struct{}{}
				}
				g[p] = struct{}{}
			}
			delete(pending, p)
		}
	}
	return
}
