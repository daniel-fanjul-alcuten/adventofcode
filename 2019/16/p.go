package p

import (
	"strconv"
	"strings"
)

func P1(s string, t int) string {
	l1, l2 := make([]byte, len(s)), make([]byte, len(s))
	for i, r := range s {
		l1[i] = byte(int(r) - '0')
	}
	for i := 0; i < t; i++ {
		for i := range l2 {
			v, j := 0, -1
			for j < len(l1) {
				j += i + 1
				k := j + i + 1
				if k > len(l1) {
					k = len(l1)
				}
				for ; j < k; j++ {
					v += int(l1[j])
				}
				j += i + 1
				k = j + i + 1
				if k > len(l1) {
					k = len(l1)
				}
				for ; j < k; j++ {
					v -= int(l1[j])
				}
			}
			if v < 0 {
				v = -v
			}
			v = v % 10
			l2[i] = byte(v)
		}
		l1, l2 = l2, l1
	}
	r := make([]rune, len(l1))
	for i := range r {
		r[i] = rune(l1[i] + '0')
	}
	return string(r)
}

func samekeys(m, n map[int]int) bool {
	for k := range m {
		if _, ok := n[k]; !ok {
			return false
		}
	}
	for k := range n {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

func P2(s string, t int) string {
	s2 := strings.Repeat(s, 10000)
	var rs func(t int, m map[int]int, n map[int]int)
	rs = func(t int, m map[int]int, n map[int]int) {
		if t == 0 {
			for i := range m {
				m[i] = int(s2[i] - '0')
			}
			return
		}
		if n == nil {
			n = map[int]int{}
			for i := range m {
				j := -1
				for j < len(s2) {
					j += i + 1
					k := j + i + 1
					if k > len(s2) {
						k = len(s2)
					}
					for ; j < k; j++ {
						n[j] = 0
					}
					j += i + 1
					k = j + i + 1
					if k > len(s2) {
						k = len(s2)
					}
					for ; j < k; j++ {
						n[j] = 0
					}
				}
			}
			if samekeys(m, n) {
				rs(t-1, n, n)
			} else {
				rs(t-1, n, nil)
			}
		} else {
			n2 := map[int]int{}
			for k := range n {
				n2[k] = 0
			}
			rs(t-1, n2, n)
			n = n2
		}
		for i := range m {
			v, j := 0, -1
			for j < len(s2) {
				j += i + 1
				k := j + i + 1
				if k > len(s2) {
					k = len(s2)
				}
				for ; j < k; j++ {
					v += n[j]
				}
				j += i + 1
				k = j + i + 1
				if k > len(s2) {
					k = len(s2)
				}
				for ; j < k; j++ {
					v -= n[j]
				}
			}
			if v < 0 {
				v = -v
			}
			v = v % 10
			m[i] = v
		}
	}
	l, err := strconv.Atoi(s2[:7])
	if err != nil {
		panic(err)
	}
	m := map[int]int{}
	for i := l; i < l+8; i++ {
		m[i] = 0
	}
	rs(t, m, nil)
	r := make([]rune, 8)
	for i := range r {
		r[i] = rune(m[l+i] + '0')
	}
	return string(r)
}
