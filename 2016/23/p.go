package p

import (
	"strconv"
	"strings"
)

type Instruction []string

func Parse(input []string) (output []Instruction) {
	output = make([]Instruction, len(input))
	for i, in := range input {
		output[i] = strings.Split(in, " ")
	}
	return
}

func P1(input []Instruction, start int) (output int) {
	i, r := 0, map[string]int{}
	r["a"] = start
	for i >= 0 && i < len(input) {
		in := input[i]
		switch in[0] {
		case "cpy":
			rx, ry := in[1], in[2]
			x, err := strconv.Atoi(rx)
			if err != nil {
				x = r[rx]
			}
			r[ry] = x
			i++
		case "inc":
			rx := in[1]
			r[rx]++
			i++
		case "dec":
			rx := in[1]
			r[rx]--
			i++
		case "jnz":
			rx, ry := in[1], in[2]
			y, err := strconv.Atoi(ry)
			if err != nil {
				y = r[ry]
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
		case "tgl":
			rx := in[1]
			x, err := strconv.Atoi(rx)
			if err != nil {
				x = r[rx]
			}
			j := i + x
			if j >= 0 && j < len(input) {
				switch input[j][0] {
				case "cpy":
					input[j][0] = "jnz"
				case "inc":
					input[j][0] = "dec"
				case "dec":
					input[j][0] = "inc"
				case "jnz":
					input[j][0] = "cpy"
				case "tgl":
					input[j][0] = "inc"
				default:
					panic(input[j][0])
				}
			}
			i++
		default:
			panic(in[0])
		}
	}
	output = r["a"]
	return
}
