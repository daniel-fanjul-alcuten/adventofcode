package p

import (
	"testing"
)

func TestP1(t *testing.T) {
	if n := P1(X); n != 7129803 {
		t.Error(n)
	}
}

func TestP2(t *testing.T) {
	t.Parallel()
	if n := P2(X); n != 12284643 {
		t.Error(n)
	}
}

var X = Program{5, []Instruction{
	{"seti", [3]int{123, 0, 4}},
	{"bani", [3]int{4, 456, 4}},
	{"eqri", [3]int{4, 72, 4}},
	{"addr", [3]int{4, 5, 5}},
	{"seti", [3]int{0, 0, 5}},
	{"seti", [3]int{0, 3, 4}},
	{"bori", [3]int{4, 65536, 1}},
	{"seti", [3]int{2024736, 3, 4}},
	{"bani", [3]int{1, 255, 2}},
	{"addr", [3]int{4, 2, 4}},
	{"bani", [3]int{4, 16777215, 4}},
	{"muli", [3]int{4, 65899, 4}},
	{"bani", [3]int{4, 16777215, 4}},
	{"gtir", [3]int{256, 1, 2}},
	{"addr", [3]int{2, 5, 5}},
	{"addi", [3]int{5, 1, 5}},
	{"seti", [3]int{27, 7, 5}},
	{"seti", [3]int{0, 1, 2}},
	{"addi", [3]int{2, 1, 3}},
	{"muli", [3]int{3, 256, 3}},
	{"gtrr", [3]int{3, 1, 3}},
	{"addr", [3]int{3, 5, 5}},
	{"addi", [3]int{5, 1, 5}},
	{"seti", [3]int{25, 2, 5}},
	{"addi", [3]int{2, 1, 2}},
	{"seti", [3]int{17, 0, 5}},
	{"setr", [3]int{2, 3, 1}},
	{"seti", [3]int{7, 9, 5}},
	{"eqrr", [3]int{4, 0, 2}},
	{"addr", [3]int{2, 5, 5}},
	{"seti", [3]int{5, 6, 5}},
}}
