package p

import (
	"testing"
)

func TestP12(t *testing.T) {
	p1, p2 := P12(X, 61, 17, 0, 1, 2)
	if p1 != 56 {
		t.Error(p1)
	}
	if p2 != 7847 {
		t.Error(p2)
	}
}
