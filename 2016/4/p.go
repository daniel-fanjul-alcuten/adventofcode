package p

import (
	"sort"
	"strconv"
	"strings"
)

type T struct {
	B rune
	C int
}

type TT []T

func (tt *TT) Len() int {
	return len(*tt)
}

func (tt *TT) Less(i, j int) bool {
	ti, tj := (*tt)[i], (*tt)[j]
	if ti.C > tj.C {
		return true
	} else if ti.C < tj.C {
		return false
	}
	return ti.B < tj.B
}

func (tt *TT) Swap(i, j int) {
	(*tt)[i], (*tt)[j] = (*tt)[j], (*tt)[i]
}

func P1(input []string) (output int) {
	for _, s := range input {
		j := strings.Index(s, "[")
		k := strings.Index(s, "]")
		i := strings.LastIndex(s, "-")
		name, sector, checksum := s[:i], s[i+1:j], s[j+1:k]
		m := make(map[rune]int, len(name))
		for _, r := range name {
			if r != '-' {
				m[r]++
			}
		}
		tt := make(TT, 0, len(m))
		for b, c := range m {
			tt = append(tt, T{b, c})
		}
		sort.Sort(&tt)
		cs := string(tt[0].B) +
			string(tt[1].B) +
			string(tt[2].B) +
			string(tt[3].B) +
			string(tt[4].B)
		if cs == checksum {
			sector, err := strconv.Atoi(sector)
			if err != nil {
				panic(err)
			}
			output += sector
		}
	}
	return
}

func P2(input []string, prefix string) (output int) {
	for _, s := range input {
		j := strings.Index(s, "[")
		k := strings.Index(s, "]")
		i := strings.LastIndex(s, "-")
		name, sector, checksum := s[:i], s[i+1:j], s[j+1:k]
		m := make(map[rune]int, len(name))
		for _, r := range name {
			if r != '-' {
				m[r]++
			}
		}
		tt := make(TT, 0, len(m))
		for b, c := range m {
			tt = append(tt, T{b, c})
		}
		sort.Sort(&tt)
		cs := string(tt[0].B) +
			string(tt[1].B) +
			string(tt[2].B) +
			string(tt[3].B) +
			string(tt[4].B)
		if cs == checksum {
			sector, err := strconv.Atoi(sector)
			if err != nil {
				panic(err)
			}
			decname := ""
			for _, r := range name {
				decname += string((int(r)-'a'+sector)%('z'-'a'+1) + 'a')
			}
			if strings.HasPrefix(decname, prefix) {
				output = sector
				return
			}
		}
	}
	return
}
