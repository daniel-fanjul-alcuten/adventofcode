package p

import (
	"strings"
)

func P1(mol string, rules [][2]string) (m map[string]struct{}) {
	m = map[string]struct{}{}
	for _, r := range rules {
		in, out := r[0], r[1]
		i := 0
		j := strings.Index(mol[i:], in)
		for j >= 0 {
			mol2 := mol[:i+j] + out + mol[i+j+len(in):]
			m[mol2] = struct{}{}
			i += j + len(in)
			j = strings.Index(mol[i:], in)
		}
	}
	return
}
