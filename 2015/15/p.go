package p

import (
	"strconv"
	"strings"
)

func P1(input []string, calories int) (output int) {
	type I struct {
		n             string
		c, d, f, t, x int
	}
	ingredients := []I{}
	for _, in := range input {
		spl := strings.Split(in, " ")
		n := strings.Trim(spl[0], ":")
		c := atoi(strings.Trim(spl[2], ","))
		d := atoi(strings.Trim(spl[4], ","))
		f := atoi(strings.Trim(spl[6], ","))
		t := atoi(strings.Trim(spl[8], ","))
		x := atoi(strings.Trim(spl[10], ","))
		g := I{n, c, d, f, t, x}
		ingredients = append(ingredients, g)
	}
	amounts := make([]int, len(ingredients))
	score := func() (int, int) {
		c, d, f, t, x := 0, 0, 0, 0, 0
		for i, n := range amounts {
			g := ingredients[i]
			c += n * g.c
			d += n * g.d
			f += n * g.f
			t += n * g.t
			x += n * g.x
		}
		if c < 0 || d < 0 || f < 0 || t < 0 {
			return 0, x
		}
		return c * d * f * t, x
	}
	var rec func(i, n int) int
	rec = func(i, n int) int {
		if i == len(amounts)-1 {
			amounts[i] = n
			sc, x := score()
			if calories != 0 && x != calories {
				return 0
			}
			return sc
		}
		sc := 0
		for m := 0; m <= n; m++ {
			amounts[i] = m
			sc2 := rec(i+1, n-m)
			if sc2 > sc {
				sc = sc2
			}
		}
		return sc
	}
	output = rec(0, 100)
	return
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
