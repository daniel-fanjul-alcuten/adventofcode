package p

import "testing"

func TestP1_S(t *testing.T) {
	if n := P1(S, 1000); n != 1120 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X, 2503); n != 2660 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if n := P2(S, 1000); n != 689 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P2(X, 2503); n != 1256 {
		t.Error(n)
	}
}

var S = []string{
	"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
	"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
}

var X = []string{
	"Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.",
	"Rudolph can fly 3 km/s for 15 seconds, but then must rest for 28 seconds.",
	"Donner can fly 19 km/s for 9 seconds, but then must rest for 164 seconds.",
	"Blitzen can fly 19 km/s for 9 seconds, but then must rest for 158 seconds.",
	"Comet can fly 13 km/s for 7 seconds, but then must rest for 82 seconds.",
	"Cupid can fly 25 km/s for 6 seconds, but then must rest for 145 seconds.",
	"Dasher can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.",
	"Dancer can fly 3 km/s for 16 seconds, but then must rest for 37 seconds.",
	"Prancer can fly 25 km/s for 6 seconds, but then must rest for 143 seconds.",
}
