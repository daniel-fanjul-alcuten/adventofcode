package p

import (
	"strconv"
	"strings"
)

func do(input []string, i int, r map[string]int) (output int) {
	for i < len(input) {
		s := strings.Split(input[i], " ")
		switch s[0] {
		case "cpy":
			rx, ry := s[1], s[2]
			x, err := strconv.Atoi(rx)
			if err != nil {
				x = r[rx]
			}
			r[ry] = x
			i++
		case "inc":
			rx := s[1]
			r[rx]++
			i++
		case "dec":
			rx := s[1]
			r[rx]--
			i++
		case "jnz":
			rx, ry := s[1], s[2]
			y, err := strconv.Atoi(ry)
			if err != nil {
				panic(y)
			}
			x, err := strconv.Atoi(rx)
			if err != nil {
				x = r[rx]
			}
			if x != 0 {
				i += y
			} else {
				i++
			}
		default:
			panic(s[0])
		}
	}
	return r["a"]
}

func P1(input []string) (output int) {
	i, r := 0, map[string]int{}
	output = do(input, i, r)
	return
}

func P2(input []string) (output int) {
	i, r := 0, map[string]int{}
	r["c"] = 1
	output = do(input, i, r)
	return
}
