package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(S); n != "tknk" {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	if n := P1(X); n != "vvsvez" {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if n := P2(S, "tknk"); n != 60 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	if n := P2(X, "vvsvez"); n != 362 {
		t.Error(n)
	}
}
