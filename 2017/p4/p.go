package p4

import (
	"sort"
	"strings"
)

func IsValid(s string) bool {
	m := map[string]struct{}{}
	for _, s := range strings.Split(s, " ") {
		if _, ok := m[s]; ok {
			return false
		}
		m[s] = struct{}{}
	}
	return true
}

func P1(ss []string) (n int) {
	for _, s := range ss {
		if IsValid(s) {
			n++
		}
	}
	return
}

type ByteArray []byte

func (a *ByteArray) Len() int {
	return len(*a)
}

func (a *ByteArray) Less(i, j int) bool {
	return (*a)[i] < (*a)[j]
}

func (a *ByteArray) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func Normalize(s string) string {
	a := ByteArray(s)
	sort.Sort(&a)
	return string(a)
}

func IsValid2(s string) bool {
	m := map[string]struct{}{}
	for _, s := range strings.Split(s, " ") {
		s = Normalize(s)
		if _, ok := m[s]; ok {
			return false
		}
		m[s] = struct{}{}
	}
	return true
}

func P2(ss []string) (n int) {
	for _, s := range ss {
		if IsValid(s) && IsValid2(s) {
			n++
		}
	}
	return
}
