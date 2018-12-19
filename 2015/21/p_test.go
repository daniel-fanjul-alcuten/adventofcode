package p

import "testing"

func TestP1(t *testing.T) {
	p := Character{100, 0, 0}
	e := Character{109, 8, 2}
	w := []Item{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}
	a := []Item{
		{"", 0, 0, 0},
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	}
	r := []Item{
		{"", 0, 0, 0},
		{"", 0, 0, 0},
		{"Damage +1", 25, 1, 0},
		{"Damage +2", 50, 2, 0},
		{"Damage +3", 100, 3, 0},
		{"Defense +1", 20, 0, 1},
		{"Defense +2", 40, 0, 2},
		{"Defense +3", 80, 0, 3},
	}
	if g := P1(p, e, w, a, r, false); g != 111 {
		t.Error(g)
	}
	if g := P1(p, e, w, a, r, true); g != 188 {
		t.Error(g)
	}
}
