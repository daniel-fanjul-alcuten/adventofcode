package p18

import (
	"testing"
)

func TestSampleP1(t *testing.T) {
	s := Execute1([]string{
		"set a 1",
		"add a 2",
		"mul a a",
		"mod a 5",
		"snd a",
		"set a 0",
		"rcv a",
		"jgz a -1",
		"set a 1",
		"jgz a -2",
	})
	if s == nil {
		t.Error(s)
	} else if *s != 4 {
		t.Error(*s)
	}
}

func TestP1(t *testing.T) {
	s := Execute1(P1)
	if s == nil {
		t.Error(s)
	} else if *s != 9423 {
		t.Error(*s)
	}
}

func TestP2(t *testing.T) {
	pos0, pos1 := 0, 0
	registers0, registers1 := map[byte]int{'p': 0}, map[byte]int{'p': 1}
	buf01, buf10 := []int{}, []int{}
	n := 0
	for {
		b1 := Execute2(P1, &pos0, registers0, &buf10, &buf01)
		l := len(buf10)
		b2 := Execute2(P1, &pos1, registers1, &buf01, &buf10)
		n += len(buf10) - l
		if !b1 && !b2 {
			break
		}
	}
	if n != 7620 {
		t.Error(n)
	}
}
