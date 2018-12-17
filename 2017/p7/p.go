package p

import (
	"strconv"
	"strings"
)

func P1(input []string) (output string) {
	d, r := map[string][]string{}, map[string][]string{}
	for _, in := range input {
		split := strings.Split(in, " ")
		m := split[0]
		if len(split) > 2 {
			for _, s := range split[3:] {
				s = strings.Trim(s, ",")
				d[m] = append(d[m], s)
				r[s] = append(r[s], m)
			}
		}
	}
	for m := range d {
		if len(r[m]) == 0 {
			return m
		}
	}
	return
}

func weight(d map[string][]string, w map[string]int, n string) (v int) {
	v += w[n]
	for _, n2 := range d[n] {
		v += weight(d, w, n2)
	}
	return
}

func P2(input []string, root string) (output int) {
	d, r, w := map[string][]string{}, map[string][]string{}, map[string]int{}
	for _, in := range input {
		split := strings.Split(in, " ")
		m := split[0]
		if len(split) > 2 {
			for _, s := range split[3:] {
				s = strings.Trim(s, ",")
				d[m] = append(d[m], s)
				r[s] = append(r[s], m)
			}
		}
		i, err := strconv.Atoi(strings.Trim(split[1], "()"))
		if err != nil {
			panic(err)
		}
		w[m] = i
	}
	n := root
	diff := 0
	for {
		w2nn := map[int][]string{}
		for _, n2 := range d[n] {
			w := weight(d, w, n2)
			w2nn[w] = append(w2nn[w], n2)
		}
		if len(w2nn) == 2 {
			for w, nn := range w2nn {
				if len(nn) != 1 {
					output = w
				}
			}
			for k, nn := range w2nn {
				if len(nn) == 1 {
					n = nn[0]
					diff = output - k
				}
			}
			continue
		}
		if len(w2nn) == 1 {
			output = w[n] + diff
			return
		}
		panic(len(w2nn))
	}
}
