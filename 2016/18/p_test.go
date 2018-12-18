package p

import (
	"testing"
)

func TestP1(t *testing.T) {
	if p, s := P1("..^^.", 1); p != ".^^^^" {
		t.Error(p)
	} else if s != 4 {
		t.Error(s)
	}
	if p, s := P1(".^^^^", 1); p != "^^..^" {
		t.Error(p)
	} else if s != 3 {
		t.Error(s)
	}
	if p, s := P1("..^^.", 2); p != "^^..^" {
		t.Error(p)
	} else if s != 6 {
		t.Error(s)
	}
	if p, s := P1(".^^.^.^^^^", 10-1); p != "^^.^^^..^^" {
		t.Error(p)
	} else if s != 38 {
		t.Error(s)
	}
	if _, s := P1(".^..^....^....^^.^^.^.^^.^.....^.^..^...^^^^^^.^^^^.^.^^^^^^^.^^^^^..^.^^^.^^..^.^^.^....^.^...^^.^.", 40-1); s != 2035 {
		t.Error(s)
	}
}

func TestP2(t *testing.T) {
	if _, s := P1(".^..^....^....^^.^^.^.^^.^.....^.^..^...^^^^^^.^^^^.^.^^^^^^^.^^^^^..^.^^^.^^..^.^^.^....^.^...^^.^.", 400000-1); s != 20000577 {
		t.Error(s)
	}
}