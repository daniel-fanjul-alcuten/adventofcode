package p23

import (
	"math"
	"strconv"
	"strings"
)

func Execute(ss []string, registers map[byte]int) (muls int) {
	pos := 0
	for pos < len(ss) {
		t := strings.Split(ss[pos], " ")
		switch t[0] {
		case "set":
			r := Register(t[1])
			v := Value(t[2], registers)
			registers[r] = v
			pos++
		case "sub":
			r := Register(t[1])
			v := Value(t[2], registers)
			registers[r] -= v
			pos++
		case "mul":
			r := Register(t[1])
			v := Value(t[2], registers)
			registers[r] *= v
			pos++
			muls++
		case "jnz":
			v := Value(t[1], registers)
			f := Value(t[2], registers)
			if v != 0 {
				pos += f
			} else {
				pos++
			}
		default:
			panic(t[0])
		}
	}
	return
}

func Register(s string) byte {
	if len(s) != 1 {
		panic(len(s))
	}
	return s[0]
}

func Value(s string, rr map[byte]int) int {
	if i, err := strconv.Atoi(s); err != nil {
		return rr[Register(s)]
	} else {
		return i
	}
}

func P2() int {
	var h int
	b := 65*100 + 100000
	c := b + 17000
	s := SieveOfEratosthenes(c)
	for {
		if s[b] {
			h += 1
		}
		if b == c {
			break
		}
		b += 17
	}
	return h
}

func SieveOfEratosthenes(value int) (s []bool) {
	s = make([]bool, value+1)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if s[i] == false {
			for j := i * i; j <= value; j += i {
				s[j] = true
			}
		}
	}
	return
}
