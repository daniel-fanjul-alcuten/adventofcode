package p16

import (
	"strconv"
	"strings"
)

type Dance []byte

func (d Dance) Spin(n int) Dance {
	i := len(d) - n
	return append(d[i:], d[:i]...)
}

func (d Dance) Exchange(i, j int) Dance {
	d2 := append(make(Dance, 0, len(d)), d...)
	d2[i], d2[j] = d2[j], d2[i]
	return d2
}

func (d Dance) Partner(a, b byte) Dance {
	d2 := make(Dance, 0, len(d))
	for _, c := range d {
		if c == a {
			d2 = append(d2, b)
		} else if c == b {
			d2 = append(d2, a)
		} else {
			d2 = append(d2, c)
		}
	}
	return d2
}

func (d Dance) Dance1(s string) Dance {
	switch s[0] {
	case 's':
		i, err := strconv.Atoi(s[1:])
		if err != nil {
			panic(err)
		}
		return d.Spin(i)
	case 'x':
		ss := strings.Split(s[1:], "/")
		i, err := strconv.Atoi(ss[0])
		if err != nil {
			panic(err)
		}
		j, err := strconv.Atoi(ss[1])
		if err != nil {
			panic(err)
		}
		return d.Exchange(i, j)
	case 'p':
		ss := strings.Split(s[1:], "/")
		return d.Partner(ss[0][0], ss[1][0])
	default:
		panic(s[0])
	}
}

func (d Dance) DanceN(s string) Dance {
	for _, s := range strings.Split(s, ",") {
		d = d.Dance1(s)
	}
	return d
}
