package p

import "log"

func P1(n int, input ...int) (g, e int) {
	for j := n - 1; j >= 0; j-- {
		n := []int{0, 0}
		for _, i := range input {
			b := (i >> j) % 2
			n[b]++
			log.Printf("%v, %05b: %v, %v", j, i, b, n)
		}
		x := 0
		if n[1] > n[0] {
			x = 1
		}
		g = 2*g + x
		e = 2*e + (1 - x)
	}
	return
}
