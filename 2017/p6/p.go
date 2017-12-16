package p6

import (
	"strconv"
	"strings"
)

type Banks []int

func (k Banks) MaxIndex() int {
	maxI, maxB := 0, 0
	for i, b := range k {
		if b > maxB {
			maxI, maxB = i, b
		}
	}
	return maxI
}

func (k Banks) Redistribute(i int) {
	n := k[i]
	k[i] = 0
	for n > 0 {
		i = (i + 1) % len(k)
		k[i] += 1
		n--
	}
}

func (k Banks) String() string {
	s := make([]string, len(k))
	for i, b := range k {
		s[i] = strconv.Itoa(b)
	}
	return strings.Join(s, ",")
}

func P1(k Banks) int {
	m := map[string]struct{}{}
	m[k.String()] = struct{}{}
	n := 0
	for {
		i := k.MaxIndex()
		k.Redistribute(i)
		n++
		s := k.String()
		if _, ok := m[s]; ok {
			return n
		}
		m[s] = struct{}{}
	}
}

func P2(k Banks) int {
	n := 0
	m := map[string]int{}
	m[k.String()] = 0
	for {
		i := k.MaxIndex()
		k.Redistribute(i)
		n++
		s := k.String()
		if m, ok := m[s]; ok {
			return n - m
		}
		m[s] = n
	}
}
