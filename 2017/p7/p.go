package p7

import (
	"strconv"
	"strings"
)

type Node struct {
	Name     string
	Weight   int
	Children []*Node
}

func Parse(ss []string) (nn map[string]*Node) {
	nn = map[string]*Node{}
	for _, s := range ss {
		p := strings.Split(s, " -> ")
		q := strings.Split(p[0], " ")
		n := Node{Name: q[0]}
		nn[n.Name] = &n
	}
	for _, s := range ss {
		p := strings.Split(s, " -> ")
		q := strings.Split(p[0], " ")
		n := nn[q[0]]
		i, j := strings.Index(q[1], "("), strings.Index(q[1], ")")
		w, err := strconv.Atoi(q[1][i+1 : j])
		if err != nil {
			panic(err)
		}
		n.Weight = w
		if len(p) > 1 {
			for _, t := range strings.Split(p[1], ", ") {
				c := nn[t]
				if c == nil {
					panic(t)
				}
				n.Children = append(n.Children, c)
			}
		}
	}
	return
}
