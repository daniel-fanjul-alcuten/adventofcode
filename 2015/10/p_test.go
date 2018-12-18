package p

import "testing"

func TestP1_S1(t *testing.T) {
	s, l := P1("211", 1)
	if s != "1221" {
		t.Error(s)
	}
	if l != 4 {
		t.Error(l)
	}
}

func TestP1_S2(t *testing.T) {
	s, l := P1("1", 1)
	if s != "11" {
		t.Error(s)
	}
	if l != 2 {
		t.Error(l)
	}
	s, l = P1("1", 2)
	if s != "21" {
		t.Error(s)
	}
	if l != 2 {
		t.Error(l)
	}
	s, l = P1("1", 3)
	if s != "1211" {
		t.Error(s)
	}
	if l != 4 {
		t.Error(l)
	}
	s, l = P1("1", 4)
	if s != "111221" {
		t.Error(s)
	}
	if l != 6 {
		t.Error(l)
	}
	s, l = P1("1", 5)
	if s != "312211" {
		t.Error(s)
	}
	if l != 6 {
		t.Error(l)
	}
}

func TestP1_X(t *testing.T) {
	_, l := P1("1113122113", 40)
	if l != 360154 {
		t.Error(l)
	}
}

func TestP2_X(t *testing.T) {
	_, l := P1("1113122113", 50)
	if l != 5103798 {
		t.Error(l)
	}
}
