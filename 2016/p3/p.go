package p

import (
	"sort"
)

func P1(input [][]int) (n int) {
	for _, s := range input {
		s := append([]int(nil), s...)
		sort.Ints(s)
		if s[0]+s[1] > s[2] {
			if s[0]+s[2] > s[1] {
				if s[1]+s[2] > s[0] {
					n++
				}
			}
		}
	}
	return
}

func P2(input [][]int) (output [][]int) {
	output = make([][]int, 0, len(input))
	for i := 0; i < len(input); i += 3 {
		output = append(output,
			[]int{input[i][0], input[i+1][0], input[i+2][0]},
			[]int{input[i][1], input[i+1][1], input[i+2][1]},
			[]int{input[i][2], input[i+1][2], input[i+2][2]},
		)
	}
	return
}
