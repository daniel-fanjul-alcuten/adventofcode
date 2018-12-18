package p

func P1(n int, input ...int) (output int) {
	if len(input) == 0 {
		return 0
	}
	output += P1(n, input[1:]...)
	if n > input[0] {
		output += P1(n-input[0], input[1:]...)
	} else if n == input[0] {
		output += 1
	}
	return
}

func P2(n int, input ...int) int {
	sols := [][]int{}
	add := func(sol []int) {
		if len(sols) == 0 || len(sol) < len(sols[0]) {
			sols = append(sols[:0], sol)
		} else if len(sol) == len(sols[0]) {
			sols = append(sols, sol)
		}
	}
	var rec func(n int, input, output []int)
	rec = func(n int, input, output []int) {
		if len(input) == 0 {
			return
		}
		rec(n, input[1:], output)
		if n > input[0] {
			rec(n-input[0], input[1:], append(output, input[0]))
		} else if n == input[0] {
			add(append(append(make([]int, 0, len(output)+1), output...), input[0]))
		}
		return
	}
	rec(n, input, nil)
	return len(sols)
}
