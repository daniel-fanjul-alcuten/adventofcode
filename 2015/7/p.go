package p

import (
	"strconv"
	"strings"
)

func P1(input []string) (output map[string]uint16) {
	m := map[string]struct{}{}
	for _, in := range input {
		m[in] = struct{}{}
	}
	output = map[string]uint16{}
	for len(m) > 0 {
		l := len(m)
		for input := range m {
			s := strings.Split(input, " ")
			switch len(s) {
			case 3:
				v, err := strconv.Atoi(s[0])
				if err != nil {
					if x, ok := output[s[0]]; ok {
						output[s[2]] = x
						delete(m, input)
					}
					continue
				}
				output[s[2]] = uint16(v)
				delete(m, input)
				continue
			case 4:
				if x, ok := output[s[1]]; ok {
					output[s[3]] = ^x
					delete(m, input)
				}
				continue
			case 5:
				switch s[1] {
				case "AND":
					x := uint16(0)
					xx, err := strconv.Atoi(s[0])
					if err == nil {
						x = uint16(xx)
					} else {
						if xx, ok := output[s[0]]; ok {
							x = xx
						} else {
							continue
						}
					}
					y := uint16(0)
					yy, err := strconv.Atoi(s[2])
					if err == nil {
						y = uint16(yy)
					} else {
						if yy, ok := output[s[2]]; ok {
							y = yy
						} else {
							continue
						}
					}
					output[s[4]] = x & y
					delete(m, input)
					continue
				case "OR":
					x := uint16(0)
					xx, err := strconv.Atoi(s[0])
					if err == nil {
						x = uint16(xx)
					} else {
						if xx, ok := output[s[0]]; ok {
							x = xx
						} else {
							continue
						}
					}
					y := uint16(0)
					yy, err := strconv.Atoi(s[2])
					if err == nil {
						y = uint16(yy)
					} else {
						if yy, ok := output[s[2]]; ok {
							y = yy
						} else {
							continue
						}
					}
					output[s[4]] = x | y
					delete(m, input)
					continue
				case "LSHIFT":
					if x, ok := output[s[0]]; ok {
						y, err := strconv.Atoi(s[2])
						if err != nil {
							panic(err)
						}
						output[s[4]] = x << uint(y)
						delete(m, input)
					}
					continue
				case "RSHIFT":
					if x, ok := output[s[0]]; ok {
						y, err := strconv.Atoi(s[2])
						if err != nil {
							panic(err)
						}
						output[s[4]] = x >> uint(y)
						delete(m, input)
					}
					continue
				}
			}
			panic(input)
		}
		if len(m) == l {
			break
		}
	}
	return
}
