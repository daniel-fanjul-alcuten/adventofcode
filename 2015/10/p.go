package p

import (
	"fmt"
	"strings"
)

func P1(input string, n int) (output string, l int) {
	output = input
	for i := 0; i < n; i++ {
		b := strings.Builder{}
		for i := 0; i < len(input); i++ {
			j := i
			for j+1 < len(input) && input[j+1] == input[j] {
				j++
			}
			if _, err := fmt.Fprintf(&b, "%d", j-i+1); err != nil {
				panic(err)
			}
			b.WriteString(input[i : i+1])
			i = j
		}
		output = b.String()
		input = output
	}
	l = len(output)
	return
}
