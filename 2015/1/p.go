package p

import "strings"

func P1(input string) (output int) {
	output += strings.Count(input, "(")
	output -= strings.Count(input, ")")
	return
}

func P2(input string) (output int) {
	f := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			f++
		} else {
			f--
		}
		if f < 0 {
			output = i + 1
			return
		}
	}
	return
}
