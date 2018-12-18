package p18

import (
	"strconv"
	"strings"
)

func Execute1(ss []string) *int {
	registers := map[byte]int{}
	lastSound := (*int)(nil)
	tt := make([][]string, 0, len(ss))
	for _, s := range ss {
		tt = append(tt, strings.Split(s, " "))
	}
	i := 0
	for {
		t := tt[i]
		switch t[0] {
		case "snd":
			v := Value(t[1], registers)
			lastSound = &v
			i++
		case "set":
			r := Register(t[1])
			v := Value(t[2], registers)
			registers[r] = v
			i++
		case "add":
			r := Register(t[1])
			v := Value(t[2], registers)
			registers[r] += v
			i++
		case "mul":
			r := Register(t[1])
			v := Value(t[2], registers)
			registers[r] *= v
			i++
		case "mod":
			r := Register(t[1])
			v := Value(t[2], registers)
			registers[r] %= v
			i++
		case "rcv":
			v := Value(t[1], registers)
			if v != 0 {
				return lastSound
			}
			i++
		case "jgz":
			v := Value(t[1], registers)
			f := Value(t[2], registers)
			if v > 0 {
				i += f
			} else {
				i++
			}
		default:
			panic(t[0])
		}
	}
}

func Execute2(ss []string, pos *int, registers map[byte]int, in *[]int, out *[]int) bool {
	t := strings.Split(ss[*pos], " ")
	switch t[0] {
	case "set":
		r := Register(t[1])
		v := Value(t[2], registers)
		registers[r] = v
		*pos++
		return true
	case "add":
		r := Register(t[1])
		v := Value(t[2], registers)
		registers[r] += v
		*pos++
		return true
	case "mul":
		r := Register(t[1])
		v := Value(t[2], registers)
		registers[r] *= v
		*pos++
		return true
	case "mod":
		r := Register(t[1])
		v := Value(t[2], registers)
		registers[r] %= v
		*pos++
		return true
	case "snd":
		v := Value(t[1], registers)
		*out = append(*out, v)
		*pos++
		return true
	case "rcv":
		if len(*in) == 0 {
			return false
		}
		r := Register(t[1])
		registers[r] = (*in)[0]
		*in = (*in)[1:]
		*pos++
		return true
	case "jgz":
		v := Value(t[1], registers)
		f := Value(t[2], registers)
		if v > 0 {
			*pos += f
		} else {
			*pos++
		}
		return true
	default:
		panic(t[0])
	}
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
