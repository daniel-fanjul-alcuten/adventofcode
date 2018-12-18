package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if m := P1(S); m != "easter" {
		t.Error(m)
	}
}

func TestP1_X(t *testing.T) {
	if m := P1(X); m != "cyxeoccr" {
		t.Error(m)
	}
}

func TestP2_S(t *testing.T) {
	if m := P2(S); m != "advent" {
		t.Error(m)
	}
}

func TestP2_X(t *testing.T) {
	if m := P2(X); m != "batwpask" {
		t.Error(m)
	}
}
