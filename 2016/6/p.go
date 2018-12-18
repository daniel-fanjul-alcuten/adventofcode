package p

import "math"

func P1(input []string) (output string) {
	count := make([]map[byte]int, len(input[0]))
	for i := range count {
		count[i] = map[byte]int{}
	}
	for _, s := range input {
		for i := 0; i < len(s); i++ {
			count[i][s[i]]++
		}
	}
	for _, m := range count {
		best, max := byte(0), 0
		for b, c := range m {
			if c > max {
				best, max = b, c
			}
		}
		output += string(best)
	}
	return
}

func P2(input []string) (output string) {
	count := make([]map[byte]int, len(input[0]))
	for i := range count {
		count[i] = map[byte]int{}
	}
	for _, s := range input {
		for i := 0; i < len(s); i++ {
			count[i][s[i]]++
		}
	}
	for _, m := range count {
		best, min := byte(0), math.MaxInt32
		for b, c := range m {
			if c < min {
				best, min = b, c
			}
		}
		output += string(best)
	}
	return
}
