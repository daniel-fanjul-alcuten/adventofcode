package p10

import (
	"testing"
)

func TestSample(t *testing.T) {
	s := New(5)
	s.ApplyN([]int{3, 4, 1, 5})
	if h := s.Hash1(); h != 12 {
		t.Error(h)
	}
}

func TestP1(t *testing.T) {
	s := New(256)
	s.ApplyN(X)
	if h := s.Hash1(); h != 3770 {
		t.Error(h)
	}
}

func TestP2(t *testing.T) {
	s := New(256)
	s.ApplyN2([]byte(Y))
	if h := s.Hash2(); h != "a9d0e68649d0174c8756a59ba21d4dc6" {
		t.Error(h)
	}
}
