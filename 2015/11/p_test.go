package p

import "testing"

func TestP1_S1(t *testing.T) {
	if s := P1("abcdefgh"); s != "abcdffaa" {
		t.Error(s)
	}
}

func TestP1_S2(t *testing.T) {
	if s := P1("ghijklmn"); s != "ghjaabcc" {
		t.Error(s)
	}
}

func TestP1_X(t *testing.T) {
	if s := P1("vzbxkghb"); s != "vzbxxyzz" {
		t.Error(s)
	}
}

func TestP2_X(t *testing.T) {
	if s := P1("vzbxxyzz"); s != "vzcaabcc" {
		t.Error(s)
	}
}
