package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(S, nil); n != 22728 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X, nil); n != 15168 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	t.Parallel()
	if n := P1(S, P2); n != 22551 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	t.Parallel()
	if n := P1(X, P2); n != 20864 {
		t.Error(n)
	}
}
