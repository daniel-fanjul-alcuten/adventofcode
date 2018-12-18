package p

import (
	"testing"
)

func TestP1(t *testing.T) {
	if s := P1("ADVENT"); s != "ADVENT" {
		t.Error(s)
	}
	if s := P1("A(1x5)BC"); s != "ABBBBBC" {
		t.Error(s)
	}
	if s := P1("(3x3)XYZ"); s != "XYZXYZXYZ" {
		t.Error(s)
	}
	if s := P1("A(2x2)BCD(2x2)EFG"); s != "ABCBCDEFEFG" {
		t.Error(s)
	}
	if s := P1("(6x1)(1x3)A"); s != "(1x3)A" {
		t.Error(s)
	}
	if s := P1("X(8x2)(3x3)ABCY"); s != "X(3x3)ABC(3x3)ABCY" {
		t.Error(s)
	}
	if n := len(P1(X)); n != 107035 {
		t.Error(n)
	}
}

func TestP2(t *testing.T) {
	if n := P2("(3x3)XYZ"); n != 9 {
		t.Error(n)
	}
	if n := P2("X(8x2)(3x3)ABCY"); n != 20 {
		t.Error(n)
	}
	if n := P2("(27x12)(20x12)(13x14)(7x10)(1x12)A"); n != 241920 {
		t.Error(n)
	}
	if n := P2("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"); n != 445 {
		t.Error(n)
	}
	if n := P2(X); n != 11451628995 {
		t.Error(n)
	}
}
