package p

import (
	"strconv"
	"strings"
)

func P1(input string) (output string) {
	for len(input) > 0 {
		i := strings.Index(input, "(")
		if i == 0 {
			j := strings.Index(input, ")")
			split := strings.Split(input[i+1:j], "x")
			l, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			n, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			input = input[j+1:]
			s := input[:l]
			input = input[len(s):]
			for ; n > 0; n-- {
				output += s
			}
		}
		if i == -1 {
			i = len(input)
		}
		s := input[:i]
		input = input[len(s):]
		output += s
	}
	return
}

func P2(input string) (output int) {
	for len(input) > 0 {
		i := strings.Index(input, "(")
		if i == 0 {
			j := strings.Index(input, ")")
			split := strings.Split(input[i+1:j], "x")
			l, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			n, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			input = input[j+1:]
			s := input[:l]
			input = input[len(s):]
			output += n * P2(s)
		}
		if i == -1 {
			i = len(input)
		}
		s := input[:i]
		input = input[len(s):]
		output += len(s)
	}
	return
}
