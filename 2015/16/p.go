package p

import (
	"strconv"
	"strings"
)

func P1(input []string, props map[string]int) (output string) {
	for _, in := range input {
		spl := strings.Split(in, " ")
		n1 := strings.Trim(spl[2], ":")
		n2 := strings.Trim(spl[4], ":")
		n3 := strings.Trim(spl[6], ":")
		m1 := atoi(strings.Trim(spl[3], ","))
		m2 := atoi(strings.Trim(spl[5], ","))
		m3 := atoi(strings.Trim(spl[7], ","))
		if props[n1] == m1 && props[n2] == m2 && props[n3] == m3 {
			output = strings.Trim(spl[1], ":")
		}
	}
	return
}

func P2(input []string, props map[string]int) (output string) {
	compare := func(n string, m int) bool {
		switch n {
		case "cats", "trees":
			return m > props[n]
		case "pomeranians", "goldfish":
			return m < props[n]
		default:
			return m == props[n]
		}
	}
	for _, in := range input {
		spl := strings.Split(in, " ")
		n1 := strings.Trim(spl[2], ":")
		n2 := strings.Trim(spl[4], ":")
		n3 := strings.Trim(spl[6], ":")
		m1 := atoi(strings.Trim(spl[3], ","))
		m2 := atoi(strings.Trim(spl[5], ","))
		m3 := atoi(strings.Trim(spl[7], ","))
		if compare(n1, m1) && compare(n2, m2) && compare(n3, m3) {
			output = strings.Trim(spl[1], ":")
		}
	}
	return
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
