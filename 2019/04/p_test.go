package p

import (
	"testing"
)

func TestP1(t *testing.T) {
	if n := P1(254032, 789860); n != 1033 {
		t.Error(n)
	}
}

func TestP2(t *testing.T) {
	if n := P2(254032, 789860); n != 670 {
		t.Error(n)
	}
}
