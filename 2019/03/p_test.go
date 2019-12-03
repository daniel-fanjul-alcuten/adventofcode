package p

import (
	"testing"
)

func TestP1(t *testing.T) {
	var w1, w2, w3 wire
	w1.fill(point{0, 0}, step{'R', 8}, step{'U', 5}, step{'L', 5}, step{'D', 3})
	w2.fill(point{0, 0}, step{'U', 7}, step{'R', 6}, step{'D', 4}, step{'L', 4})
	w3.intersect(w1, w2)
	if d := w3.distance1(); d != 6 {
		t.Error(d)
	}
	w1.fill(point{0, 0}, S1[0]...)
	w2.fill(point{0, 0}, S1[1]...)
	w3.intersect(w1, w2)
	if d := w3.distance1(); d != 159 {
		t.Error(d)
	}
	w1.fill(point{0, 0}, S2[0]...)
	w2.fill(point{0, 0}, S2[1]...)
	w3.intersect(w1, w2)
	if d := w3.distance1(); d != 135 {
		t.Error(d)
	}
	w1.fill(point{0, 0}, X1...)
	w2.fill(point{0, 0}, X2...)
	w3.intersect(w1, w2)
	if d := w3.distance1(); d != 232 {
		t.Error(d)
	}
}

func TestP2(t *testing.T) {
	var w1, w2, w3 wire
	w1.fill(point{0, 0}, step{'R', 8}, step{'U', 5}, step{'L', 5}, step{'D', 3})
	w2.fill(point{0, 0}, step{'U', 7}, step{'R', 6}, step{'D', 4}, step{'L', 4})
	w3.intersect(w1, w2)
	if d := w3.distance2(); d != 30 {
		t.Error(d)
	}
	w1.fill(point{0, 0}, S1[0]...)
	w2.fill(point{0, 0}, S1[1]...)
	w3.intersect(w1, w2)
	if d := w3.distance2(); d != 610 {
		t.Error(d)
	}
	w1.fill(point{0, 0}, S2[0]...)
	w2.fill(point{0, 0}, S2[1]...)
	w3.intersect(w1, w2)
	if d := w3.distance2(); d != 410 {
		t.Error(d)
	}
	w1.fill(point{0, 0}, X1...)
	w2.fill(point{0, 0}, X2...)
	w3.intersect(w1, w2)
	if d := w3.distance2(); d != 6084 {
		t.Error(d)
	}
}
