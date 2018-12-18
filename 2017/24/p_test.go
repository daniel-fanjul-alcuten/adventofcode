package p24

import (
	"testing"
)

func TestS1(t *testing.T) {
	r := Parse(S1)
	if s := Strongest(0, 0, r); s != 31 {
		t.Error(s)
	}
}

func TestP1(t *testing.T) {
	r := Parse(P1)
	if s := Strongest(0, 0, r); s != 1656 {
		t.Error(s)
	}
}

func TestS2(t *testing.T) {
	r := Parse(S1)
	l, s := LongestStrongest(0, 0, 0, r)
	if l != 4 {
		t.Error(l)
	}
	if s != 19 {
		t.Error(s)
	}
}

func TestP2(t *testing.T) {
	r := Parse(P1)
	l, s := LongestStrongest(0, 0, 0, r)
	if l != 30 {
		t.Error(l)
	}
	if s != 1642 {
		t.Error(s)
	}
}
