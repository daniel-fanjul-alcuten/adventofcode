package p

import "testing"

func TestP1_S(t *testing.T) {
	if n := P1(S, 0); n != 62842880 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X, 0); n != 222870 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if n := P1(S, 500); n != 57600000 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P1(X, 500); n != 117936 {
		t.Error(n)
	}
}

var S = []string{
	"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
	"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
}

var X = []string{
	"Sugar: capacity 3, durability 0, flavor 0, texture -3, calories 2",
	"Sprinkles: capacity -3, durability 3, flavor 0, texture 0, calories 9",
	"Candy: capacity -1, durability 0, flavor 4, texture 0, calories 1",
	"Chocolate: capacity 0, durability 0, flavor -2, texture 2, calories 8",
}
