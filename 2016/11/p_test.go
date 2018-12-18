package p

import (
	"testing"
)

func TestP2_S(t *testing.T) {
	t.Parallel()
	if n := S.P2(); n != 11 {
		t.Error(n)
	}
}

func TestP2_X1(t *testing.T) {
	t.Parallel()
	if n := X1.P2(); n != 37 {
		t.Error(n)
	}
}

func TestP2_X2(t *testing.T) {
	t.Parallel()
	if n := X2.P2(); n != 61 {
		t.Error(n)
	}
}
