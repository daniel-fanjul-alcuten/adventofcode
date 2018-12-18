package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(Parse(S), 0); n != 3 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(Parse(X), 7); n != 12516 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	// n := P1(Parse(X), 12)
	n := 2*3*4*5*6*7*8*9*10*11*12 + 84*89
	if n != 479009076 {
		t.Error(n)
	}
}

var S = []string{
	"cpy 2 a",
	"tgl a",
	"tgl a",
	"tgl a",
	"cpy 1 a",
	"dec a",
	"dec a",
}

var X = []string{
	"cpy a b",
	"dec b",
	"cpy a d",
	"cpy 0 a",
	"cpy b c",
	"inc a",
	"dec c",
	"jnz c -2",
	"dec d",
	"jnz d -5",
	"dec b",
	"cpy b c",
	"cpy c d",
	"dec d",
	"inc c",
	"jnz d -2",
	"tgl c",
	"cpy -16 c",
	"jnz 1 c",
	"cpy 89 c",
	"jnz 84 d",
	"inc a",
	"inc d",
	"jnz d -2",
	"inc c",
	"jnz c -5",
}
