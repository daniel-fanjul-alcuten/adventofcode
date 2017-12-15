package p2

import (
	"math"
	"strconv"
	"strings"
)

func Parse1(s string) (r []int) {
	for _, s := range strings.Split(s, " ") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		r = append(r, i)
	}
	return
}

func ParseN(ss []string) (r [][]int) {
	for _, s := range ss {
		r = append(r, Parse1(s))
	}
	return
}

func P1(ppp [][]int) (r int) {

	for _, pp := range ppp {
		min, max := math.MaxInt32, math.MinInt32
		for _, p := range pp {
			if p < min {
				min = p
			}
			if p > max {
				max = p
			}
		}
		r += max - min
	}
	return
}

func P2(ppp [][]int) (r int) {

	for _, pp := range ppp {
		m := 0
	L:
		for i, p1 := range pp {
			for j, p2 := range pp {
				if i != j {
					if p1%p2 == 0 {
						m = p1 / p2
						break L
					}
					if p2%p1 == 0 {
						m = p2 / p1
						break L
					}
				}
			}
		}
		r += m
	}
	return
}
