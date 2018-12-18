package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if n := P1(S); n != 5 {
		t.Error(n)
	}
}

func TestP1_X1(t *testing.T) {
	if n := P1(X1); n != 376777 {
		t.Error(n)
	}
}

func TestP1_X2(t *testing.T) {
	if n := P1(X2); n != 3903937 {
		t.Error(n)
	}
}

func TestP3_S(t *testing.T) {
	if n := P3(S); n != 5 {
		t.Error(n)
	}
}

func TestP3_X1(t *testing.T) {
	if n := P3(X1); n != 376777 {
		t.Error(n)
	}
	if b := Check(X1, 376777); !b {
		t.Error(b)
	}
}

func TestP3_X2(t *testing.T) {
	if n := P3(X2); n != 3903937 {
		t.Error(n)
	}
	if b := Check(X2, 3903937); !b {
		t.Error(b)
	}
}

func TestP3_X3(t *testing.T) {
	if n := P3(X3); n != 213571335378946582 {
		t.Error(n)
	}
	if b := Check(X3, 213571335378946582); !b {
		t.Error(b)
	}
}
