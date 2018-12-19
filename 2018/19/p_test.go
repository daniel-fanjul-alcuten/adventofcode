package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(S, 0); n != 6 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X, 0); n != 960 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	r0, r4 := 0, 10551293
	for r2 := 1; r2 <= r4; r2++ {
		if r4%r2 == 0 {
			r0 += r2
		}
	}
	if r0 != 10750428 {
		t.Fatal(r0)
	}
	return
}

var S = Program{0, []Instruction{
	{"seti", [3]int{5, 0, 1}},
	{"seti", [3]int{6, 0, 2}},
	{"addi", [3]int{0, 1, 0}},
	{"addr", [3]int{1, 2, 3}},
	{"setr", [3]int{1, 0, 0}},
	{"seti", [3]int{8, 0, 4}},
	{"seti", [3]int{9, 0, 5}},
}}

var X = Program{5, []Instruction{
	{"addi", [3]int{5, 16, 5}},
	{"seti", [3]int{1, 1, 2}},
	{"seti", [3]int{1, 8, 1}},
	{"mulr", [3]int{2, 1, 3}},
	{"eqrr", [3]int{3, 4, 3}},
	{"addr", [3]int{3, 5, 5}},
	{"addi", [3]int{5, 1, 5}},
	{"addr", [3]int{2, 0, 0}},
	{"addi", [3]int{1, 1, 1}},
	{"gtrr", [3]int{1, 4, 3}},
	{"addr", [3]int{5, 3, 5}},
	{"seti", [3]int{2, 6, 5}},
	{"addi", [3]int{2, 1, 2}},
	{"gtrr", [3]int{2, 4, 3}},
	{"addr", [3]int{3, 5, 5}},
	{"seti", [3]int{1, 2, 5}},
	{"mulr", [3]int{5, 5, 5}},
	{"addi", [3]int{4, 2, 4}},
	{"mulr", [3]int{4, 4, 4}},
	{"mulr", [3]int{5, 4, 4}},
	{"muli", [3]int{4, 11, 4}},
	{"addi", [3]int{3, 2, 3}},
	{"mulr", [3]int{3, 5, 3}},
	{"addi", [3]int{3, 13, 3}},
	{"addr", [3]int{4, 3, 4}},
	{"addr", [3]int{5, 0, 5}},
	{"seti", [3]int{0, 8, 5}},
	{"setr", [3]int{5, 5, 3}},
	{"mulr", [3]int{3, 5, 3}},
	{"addr", [3]int{5, 3, 3}},
	{"mulr", [3]int{5, 3, 3}},
	{"muli", [3]int{3, 14, 3}},
	{"mulr", [3]int{3, 5, 3}},
	{"addr", [3]int{4, 3, 4}},
	{"seti", [3]int{0, 9, 0}},
	{"seti", [3]int{0, 9, 5}},
}}
