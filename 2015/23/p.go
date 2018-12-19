package p

import (
	"strconv"
	"strings"
)

type Operation func(pc0, a0, b0 int) (pc1, a1, b1 int)

func P1(input []string, a0 int) (a, b int) {
	ops := make([]Operation, len(input))
	for i, in := range input {
		spl := strings.Split(in, " ")
		switch spl[0] {
		case "hlf":
			switch spl[1] {
			case "a":
				ops[i] = func(pc, a, b int) (int, int, int) {
					pc++
					a /= 2
					return pc, a, b
				}
			default:
				panic(spl[1])
			}
		case "tpl":
			switch spl[1] {
			case "a":
				ops[i] = func(pc, a, b int) (int, int, int) {
					pc++
					a *= 3
					return pc, a, b
				}
			default:
				panic(spl[1])
			}
		case "inc":
			switch spl[1] {
			case "a":
				ops[i] = func(pc, a, b int) (int, int, int) {
					pc++
					a++
					return pc, a, b
				}
			case "b":
				ops[i] = func(pc, a, b int) (int, int, int) {
					pc++
					b++
					return pc, a, b
				}
			default:
				panic(spl[1])
			}
		case "jmp":
			f, err := strconv.Atoi(spl[1])
			if err != nil {
				panic(err)
			}
			ops[i] = func(pc, a, b int) (int, int, int) {
				pc += f
				return pc, a, b
			}
		case "jie":
			f, err := strconv.Atoi(spl[2])
			if err != nil {
				panic(err)
			}
			switch strings.Trim(spl[1], ",") {
			case "a":
				ops[i] = func(pc, a, b int) (int, int, int) {
					if a%2 == 0 {
						pc += f
					} else {
						pc++
					}
					return pc, a, b
				}
			default:
				panic(spl[1])
			}
		case "jio":
			f, err := strconv.Atoi(spl[2])
			if err != nil {
				panic(err)
			}
			switch strings.Trim(spl[1], ",") {
			case "a":
				ops[i] = func(pc, a, b int) (int, int, int) {
					if a == 1 {
						pc += f
					} else {
						pc++
					}
					return pc, a, b
				}
			default:
				panic(spl[1])
			}
		default:
			panic(spl[0])
		}
	}
	pc, a, b := 0, a0, 0
	for pc >= 0 && pc < len(ops) {
		pc, a, b = ops[pc](pc, a, b)
	}
	return
}
