package p

import (
	"fmt"
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

func Execute(input []Instruction, a int) (output bool) {
	type E struct {
		V int
		T int
	}
	i, r, e := 0, map[string]int{}, []E{}
	r["a"] = a
	state := func() string {
		return fmt.Sprintf("%v:%v,%v,%v,%v", i, r["a"], r["b"], r["c"], r["d"])
	}
	states := map[string]int{state(): 0}
	for l := 1; i >= 0 && i < len(input); l++ {
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
		case "out":
			rx := in[1]
			x, err := strconv.Atoi(rx)
			if err != nil {
				x = r[rx]
			}
			e = append(e, E{x, l})
			if l := len(e); l > 1 {
				if e[l-2].V == e[l-1].V {
					return false
				}
			}
			i++
		default:
			panic(in[0])
		}
		s := state()
		if l2, ok := states[s]; ok {
			if len(e) == 0 {
				return false
			}
			for _, c := range e {
				if c.T > l2 {
					return c.V != e[len(e)-1].V
				}
			}
			return false
		}
		states[s] = l
	}
	return false
}

func P1(input []Instruction) int {
	for a := 1; ; a++ {
		if Execute(input, a) {
			return a
		}
	}
}
