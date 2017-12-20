package p19

import (
	"testing"
)

func TestS1(t *testing.T) {
	if p, n := Traverse(S1); p != "ABCDEF" {
		t.Error(p)
	} else if n != 38 {
		t.Error(n)
	}
}

func TestP1(t *testing.T) {
	if p, n := Traverse(P1); p != "HATBMQJYZ" {
		t.Error(p)
	} else if n != 16332 {
		t.Error(n)
	}
}
