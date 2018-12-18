package p

import (
	"strconv"
	"strings"
)

func P1(input []string) (output int) {
	type P struct {
		X, Y int
	}
	parseP := func(s string) P {
		p := strings.Split(s, ",")
		x, err := strconv.Atoi(p[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(p[1])
		if err != nil {
			panic(err)
		}
		return P{x, y}
	}
	l := map[P]struct{}{}
	for _, in := range input {
		s := strings.Split(in, " ")
		switch s[0] {
		case "turn":
			ok := s[1] != "on"
			p0, p1 := parseP(s[2]), parseP(s[4])
			for y := p0.Y; y <= p1.Y; y++ {
				for x := p0.X; x <= p1.X; x++ {
					p := P{x, y}
					if ok {
						delete(l, p)
					} else {
						l[p] = struct{}{}
					}
				}
			}
		case "toggle":
			p0, p1 := parseP(s[1]), parseP(s[3])
			for y := p0.Y; y <= p1.Y; y++ {
				for x := p0.X; x <= p1.X; x++ {
					p := P{x, y}
					if _, ok := l[p]; ok {
						delete(l, p)
					} else {
						l[p] = struct{}{}
					}
				}
			}
		default:
			panic(s[0])
		}
	}
	output = len(l)
	return
}

func P2(input []string) (output int) {
	type P struct {
		X, Y int
	}
	parseP := func(s string) P {
		p := strings.Split(s, ",")
		x, err := strconv.Atoi(p[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(p[1])
		if err != nil {
			panic(err)
		}
		return P{x, y}
	}
	l := map[P]int{}
	for _, in := range input {
		s := strings.Split(in, " ")
		switch s[0] {
		case "turn":
			on := s[1] == "on"
			p0, p1 := parseP(s[2]), parseP(s[4])
			for y := p0.Y; y <= p1.Y; y++ {
				for x := p0.X; x <= p1.X; x++ {
					p := P{x, y}
					if on {
						l[p]++
					} else if l[p] > 0 {
						l[p]--
					}
				}
			}
		case "toggle":
			p0, p1 := parseP(s[1]), parseP(s[3])
			for y := p0.Y; y <= p1.Y; y++ {
				for x := p0.X; x <= p1.X; x++ {
					p := P{x, y}
					l[p] += 2
				}
			}
		default:
			panic(s[0])
		}
	}
	for _, l := range l {
		output += l
	}
	return
}
