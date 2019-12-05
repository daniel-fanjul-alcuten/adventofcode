package p

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	s := []int{3, 0, 4, 0, 99}
	if outputs, err := Intcode(12345, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[12345]" {
		t.Error(outputs)
	}
	if s := fmt.Sprint(s); s != "[12345 0 4 0 99]" {
		t.Error(s)
	}
	s = []int{1002, 4, 3, 4, 33}
	if outputs, err := Intcode(12345, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[]" {
		t.Error(outputs)
	}
	if s := fmt.Sprint(s); s != "[1002 4 3 4 99]" {
		t.Error(s)
	}
	s = []int{1101, 100, -1, 4, 0}
	if outputs, err := Intcode(12345, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[]" {
		t.Error(outputs)
	}
	if s := fmt.Sprint(s); s != "[1101 100 -1 4 99]" {
		t.Error(s)
	}
	s = append(make([]int, 0, len(X)), X...)
	if outputs, err := Intcode(1, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0 0 0 0 0 0 0 0 0 16225258]" {
		t.Error(outputs)
	}
}

func TestP2(t *testing.T) {
	s := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	if outputs, err := Intcode(7, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	if outputs, err := Intcode(8, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[1]" {
		t.Error(outputs)
	}
	s = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	if outputs, err := Intcode(9, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	if outputs, err := Intcode(7, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[1]" {
		t.Error(outputs)
	}
	s = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	if outputs, err := Intcode(8, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	if outputs, err := Intcode(9, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	if outputs, err := Intcode(7, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	if outputs, err := Intcode(8, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[1]" {
		t.Error(outputs)
	}
	s = []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	if outputs, err := Intcode(9, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	if outputs, err := Intcode(7, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[1]" {
		t.Error(outputs)
	}
	s = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	if outputs, err := Intcode(8, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	if outputs, err := Intcode(9, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	if outputs, err := Intcode(0, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	if outputs, err := Intcode(1, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[1]" {
		t.Error(outputs)
	}
	s = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	if outputs, err := Intcode(0, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[0]" {
		t.Error(outputs)
	}
	s = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	if outputs, err := Intcode(1, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[1]" {
		t.Error(outputs)
	}
	s = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999,
		1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	if outputs, err := Intcode(7, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[999]" {
		t.Error(outputs)
	}
	s = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999,
		1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	if outputs, err := Intcode(8, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[1000]" {
		t.Error(outputs)
	}
	s = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999,
		1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	if outputs, err := Intcode(9, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[1001]" {
		t.Error(outputs)
	}
	s = append(make([]int, 0, len(X)), X...)
	if outputs, err := Intcode(5, &s); err != nil {
		t.Error(err)
	} else if s := fmt.Sprint(outputs); s != "[2808771]" {
		t.Error(outputs)
	}
}

var X = []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1101, 90, 60,
	224, 1001, 224, -150, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1,
	224, 223, 223, 1, 57, 83, 224, 1001, 224, -99, 224, 4, 224, 1002, 223, 8,
	223, 1001, 224, 5, 224, 1, 223, 224, 223, 1102, 92, 88, 225, 101, 41, 187,
	224, 1001, 224, -82, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1,
	224, 223, 223, 1101, 7, 20, 225, 1101, 82, 64, 225, 1002, 183, 42, 224, 101,
	-1554, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223,
	223, 1102, 70, 30, 224, 101, -2100, 224, 224, 4, 224, 102, 8, 223, 223, 101,
	1, 224, 224, 1, 224, 223, 223, 2, 87, 214, 224, 1001, 224, -2460, 224, 4,
	224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 223, 224, 223, 102, 36, 180,
	224, 1001, 224, -1368, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 5, 224, 1,
	223, 224, 223, 1102, 50, 38, 225, 1102, 37, 14, 225, 1101, 41, 20, 225, 1001,
	217, 7, 224, 101, -25, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 2, 224, 224,
	1, 224, 223, 223, 1101, 7, 30, 225, 1102, 18, 16, 225, 4, 223, 99, 0, 0, 0,
	677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105,
	1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999,
	1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999,
	1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0,
	1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0,
	0, 106, 0, 0, 1105, 1, 99999, 7, 226, 226, 224, 102, 2, 223, 223, 1006, 224,
	329, 101, 1, 223, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 344,
	1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 359, 101,
	1, 223, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 374, 101, 1,
	223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 389, 101, 1, 223,
	223, 108, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 404, 101, 1, 223, 223,
	1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 419, 101, 1, 223, 223, 8,
	226, 677, 224, 102, 2, 223, 223, 1006, 224, 434, 1001, 223, 1, 223, 1008,
	677, 677, 224, 1002, 223, 2, 223, 1005, 224, 449, 1001, 223, 1, 223, 1107,
	226, 677, 224, 102, 2, 223, 223, 1006, 224, 464, 101, 1, 223, 223, 107, 226,
	677, 224, 1002, 223, 2, 223, 1006, 224, 479, 1001, 223, 1, 223, 7, 226, 677,
	224, 102, 2, 223, 223, 1005, 224, 494, 1001, 223, 1, 223, 8, 677, 677, 224,
	102, 2, 223, 223, 1006, 224, 509, 1001, 223, 1, 223, 1108, 677, 677, 224,
	102, 2, 223, 223, 1005, 224, 524, 1001, 223, 1, 223, 1108, 226, 677, 224,
	1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 107, 226, 226, 224, 102,
	2, 223, 223, 1006, 224, 554, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2,
	223, 223, 1005, 224, 569, 1001, 223, 1, 223, 1008, 226, 226, 224, 102, 2,
	223, 223, 1005, 224, 584, 101, 1, 223, 223, 1007, 677, 677, 224, 1002, 223,
	2, 223, 1005, 224, 599, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2,
	223, 1006, 224, 614, 1001, 223, 1, 223, 1007, 226, 677, 224, 1002, 223, 2,
	223, 1006, 224, 629, 101, 1, 223, 223, 1008, 677, 226, 224, 102, 2, 223, 223,
	1005, 224, 644, 101, 1, 223, 223, 1107, 226, 226, 224, 1002, 223, 2, 223,
	1005, 224, 659, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223,
	1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
