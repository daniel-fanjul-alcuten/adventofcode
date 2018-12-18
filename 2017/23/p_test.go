package p23

import (
	"testing"
)

func TestP1(t *testing.T) {
	registers := map[byte]int{}
	muls := Execute(P1, registers)
	if muls != 3969 {
		t.Error(muls)
	}
}

func TestP2(t *testing.T) {
	if h := P2(); h != 917 {
		t.Error(h)
	}
}
