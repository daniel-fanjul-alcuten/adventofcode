package p

import (
	"testing"
)

func TestP12_S(t *testing.T) {
	p1, p2 := P12(S)
	if p1 != 57 {
		t.Error(p1)
	}
	if p2 != 29 {
		t.Error(p2)
	}
}

func TestP12_X(t *testing.T) {
	p1, p2 := P12(X)
	if p1 != 44743 {
		t.Error(p1)
	}
	if p2 != 34172 {
		t.Error(p2)
	}
}
