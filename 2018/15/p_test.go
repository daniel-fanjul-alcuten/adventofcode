package p

import (
	"testing"
)

func TestP1_S1(t *testing.T) {
	z := New(3, S1)
	if n := z.P1(); n != 47*590 {
		t.Error(n)
	}
}

func TestP1_S2(t *testing.T) {
	z := New(3, S2)
	if n := z.P1(); n != 37*982 {
		t.Error(n)
	}
}

func TestP1_S3(t *testing.T) {
	z := New(3, S3)
	if n := z.P1(); n != 46*859 {
		t.Error(n)
	}
}

func TestP1_S4(t *testing.T) {
	z := New(3, S4)
	if n := z.P1(); n != 35*793 {
		t.Error(n)
	}
}

func TestP1_S5(t *testing.T) {
	z := New(3, S5)
	if n := z.P1(); n != 54*536 {
		t.Error(n)
	}
}

func TestP1_S6(t *testing.T) {
	z := New(3, S6)
	if n := z.P1(); n != 20*937 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	z := New(3, X)
	if n := z.P1(); n != 64*2812 {
		t.Error(n)
	}
}

func TestP2_S1(t *testing.T) {
	z := New(3, S1)
	if n := z.P2(); n != 29*172 {
		t.Error(n)
	}
}

func TestP2_S2(t *testing.T) {
	z := New(3, S2)
	if n := z.P2(); n != 28*1038 {
		t.Error(n)
	}
}

func TestP2_S3(t *testing.T) {
	z := New(3, S3)
	if n := z.P2(); n != 33*948 {
		t.Error(n)
	}
}

func TestP2_S4(t *testing.T) {
	z := New(3, S4)
	if n := z.P2(); n != 37*94 {
		t.Error(n)
	}
}

func TestP2_S5(t *testing.T) {
	z := New(3, S5)
	if n := z.P2(); n != 39*166 {
		t.Error(n)
	}
}

func TestP2_S6(t *testing.T) {
	z := New(3, S6)
	if n := z.P2(); n != 30*38 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	z := New(3, X)
	if n := z.P2(); n != 31*1358 {
		t.Error(n)
	}
}
