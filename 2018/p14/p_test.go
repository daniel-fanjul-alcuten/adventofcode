package p

import (
	"testing"
)

func TestP1_5(t *testing.T) {
	if s := P1(5); s != "0124515891" {
		t.Error(s)
	}
}

func TestP1_9(t *testing.T) {
	if s := P1(9); s != "5158916779" {
		t.Error(s)
	}
}

func TestP1_18(t *testing.T) {
	if s := P1(18); s != "9251071085" {
		t.Error(s)
	}
}

func TestP1_2018(t *testing.T) {
	if s := P1(2018); s != "5941429882" {
		t.Error(s)
	}
}

func TestP1_X(t *testing.T) {
	if s := P1(157901); s != "9411137133" {
		t.Error(s)
	}
}

func TestP2_51589(t *testing.T) {
	if n := P2(51589, "51589"); n != 9 {
		t.Error(n)
	}
}

func TestP2_01245(t *testing.T) {
	if n := P2(1245, "01245"); n != 5 {
		t.Error(n)
	}
}

func TestP2_92510(t *testing.T) {
	if n := P2(92510, "92510"); n != 18 {
		t.Error(n)
	}
}

func TestP2_59414(t *testing.T) {
	if n := P2(59414, "59414"); n != 2018 {
		t.Error(n)
	}
}

func TestP2_157901(t *testing.T) {
	if n := P2(157901, "157901"); n != 20317612 {
		t.Error(n)
	}
}
