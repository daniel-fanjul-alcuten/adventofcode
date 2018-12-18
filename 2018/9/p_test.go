package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if s := P1(9, 25); s != 32 {
		t.Error(s)
	}
	if s := P1(10, 1618); s != 8317 {
		t.Error(s)
	}
	if s := P1(13, 7999); s != 146373 {
		t.Error(s)
	}
	if s := P1(17, 1104); s != 2764 {
		t.Error(s)
	}
	if s := P1(21, 6111); s != 54718 {
		t.Error(s)
	}
	if s := P1(30, 5807); s != 37305 {
		t.Error(s)
	}
}

func TestP1_X(t *testing.T) {
	if s := P1(473, 70904); s != 371284 {
		t.Error(s)
	}
}

func TestP2_X(t *testing.T) {
	if s := P1(473, 100*70904); s != 3038972494 {
		t.Error(s)
	}
}
