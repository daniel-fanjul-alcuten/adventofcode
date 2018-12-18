package p

import (
	"bytes"
	"strconv"
	"strings"
)

func P1(s string, rr []string) (m string) {
	p := []byte(s)
	swapPosition := func(x, y int) {
		p[x], p[y] = p[y], p[x]
	}
	index := func(x byte) int {
		for i, b := range p {
			if b == x {
				return i
			}
		}
		panic(x)
	}
	swapLetter := func(x, y byte) {
		swapPosition(index(x), index(y))
	}
	rotateLeft := func(x int) {
		for i := 0; i < x; i++ {
			p = append(p[1:], p[0])
		}
	}
	rotateRight := func(x int) {
		for i := 0; i < x; i++ {
			l := len(p)
			p = append(p[l-1:l], p[:l-1]...)
		}
	}
	rotateBased := func(x byte) {
		i := index(x)
		r := 1 + i
		if i >= 4 {
			r++
		}
		rotateRight(r)
	}
	reversePositions := func(x, y int) {
		for x < y {
			swapPosition(x, y)
			x, y = x+1, y-1
		}
	}
	movePositions := func(x, y int) {
		b := p[x]
		p = append(p[:x], p[x+1:]...)
		p = append(p, b)
		copy(p[y+1:], p[y:])
		p[y] = b
	}
	for _, r := range rr {
		s := strings.Split(r, " ")
		switch s[0] {
		case "swap":
			switch s[1] {
			case "position":
				x, err := strconv.Atoi(s[2])
				if err != nil {
					panic(err)
				}
				y, err := strconv.Atoi(s[5])
				if err != nil {
					panic(err)
				}
				swapPosition(x, y)
			case "letter":
				x := s[2][0]
				y := s[5][0]
				swapLetter(x, y)
			default:
				panic(s[1])
			}
		case "rotate":
			switch s[1] {
			case "left":
				x, err := strconv.Atoi(s[2])
				if err != nil {
					panic(err)
				}
				rotateLeft(x)
			case "right":
				x, err := strconv.Atoi(s[2])
				if err != nil {
					panic(err)
				}
				rotateRight(x)
			case "based":
				x := s[6][0]
				rotateBased(x)
			default:
				panic(s[1])
			}
		case "reverse":
			x, err := strconv.Atoi(s[2])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(s[4])
			if err != nil {
				panic(err)
			}
			reversePositions(x, y)
		case "move":
			x, err := strconv.Atoi(s[2])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(s[5])
			if err != nil {
				panic(err)
			}
			movePositions(x, y)
		default:
			panic(s[0])
		}
	}
	m = string(p)
	return
}

func P2(s string, rr []string) (m string) {
	p := []byte(s)
	swapPosition := func(x, y int) {
		p[x], p[y] = p[y], p[x]
	}
	unswapPosition := swapPosition
	index := func(x byte) int {
		for i, b := range p {
			if b == x {
				return i
			}
		}
		panic(x)
	}
	swapLetter := func(x, y byte) {
		swapPosition(index(x), index(y))
	}
	unswapLetter := swapLetter
	rotateLeft := func(x int) {
		for i := 0; i < x; i++ {
			p = append(p[1:], p[0])
		}
	}
	rotateRight := func(x int) {
		for i := 0; i < x; i++ {
			l := len(p)
			p = append(p[l-1:l], p[:l-1]...)
		}
	}
	unrotateLeft := rotateRight
	unrotateRight := rotateLeft
	rotateBased := func(x byte) {
		i := index(x)
		r := 1 + i
		if i >= 4 {
			r++
		}
		rotateRight(r)
	}
	unrotateBased := func(x byte) {
		p0 := append(make([]byte, 0, len(p)), p...)
		q := make([]byte, len(p))
		for {
			rotateLeft(1)
			copy(q, p)
			rotateBased(x)
			if bytes.Compare(p, p0) == 0 {
				copy(p, q)
				return
			}
			copy(p, q)
		}
	}
	reversePositions := func(x, y int) {
		for x < y {
			swapPosition(x, y)
			x, y = x+1, y-1
		}
	}
	unreversePositions := reversePositions
	movePositions := func(x, y int) {
		b := p[x]
		p = append(p[:x], p[x+1:]...)
		p = append(p, b)
		copy(p[y+1:], p[y:])
		p[y] = b
	}
	unmovePositions := func(x, y int) {
		movePositions(y, x)
	}
	for i := len(rr) - 1; i >= 0; i-- {
		r := rr[i]
		s := strings.Split(r, " ")
		switch s[0] {
		case "swap":
			switch s[1] {
			case "position":
				x, err := strconv.Atoi(s[2])
				if err != nil {
					panic(err)
				}
				y, err := strconv.Atoi(s[5])
				if err != nil {
					panic(err)
				}
				unswapPosition(x, y)
			case "letter":
				x := s[2][0]
				y := s[5][0]
				unswapLetter(x, y)
			default:
				panic(s[1])
			}
		case "rotate":
			switch s[1] {
			case "left":
				x, err := strconv.Atoi(s[2])
				if err != nil {
					panic(err)
				}
				unrotateLeft(x)
			case "right":
				x, err := strconv.Atoi(s[2])
				if err != nil {
					panic(err)
				}
				unrotateRight(x)
			case "based":
				x := s[6][0]
				unrotateBased(x)
			default:
				panic(s[1])
			}
		case "reverse":
			x, err := strconv.Atoi(s[2])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(s[4])
			if err != nil {
				panic(err)
			}
			unreversePositions(x, y)
		case "move":
			x, err := strconv.Atoi(s[2])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(s[5])
			if err != nil {
				panic(err)
			}
			unmovePositions(x, y)
		default:
			panic(s[0])
		}
	}
	m = string(p)
	return
}
