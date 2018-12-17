package p

import (
	"container/ring"
)

func P1(input int) (output int) {
	r := ring.New(input)
	for i := 0; i < input; i++ {
		r.Value = i + 1
		r = r.Next()
	}
	for i := 1; i < input; i++ {
		r.Unlink(1)
		r = r.Next()
	}
	output = r.Value.(int)
	return
}

func P2(input int) (output int) {
	r := ring.New(input)
	for i := 0; i < input; i++ {
		r.Value = i + 1
		r = r.Next()
	}
	s := r.Move(input/2 - 1)
	for n := input; n > 1; n-- {
		s.Unlink(1)
		r = r.Next()
		if n%2 == 1 {
			s = s.Next()
		}
	}
	output = r.Value.(int)
	return
}
