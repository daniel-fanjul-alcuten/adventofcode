package p

import (
	"testing"
)

func TestP1_S1_1(t *testing.T) {
	if p := P1(S1, 1); p != nil {
		t.Error(p)
	}
}

func TestP1_S1_2(t *testing.T) {
	if p := P1(S1, 2); p == nil {
		t.Error(p)
	} else if *p != (Position{0, 3}) {
		t.Error(*p)
	}
}

func TestP1_S2_13(t *testing.T) {
	if p := P1(S2, 13); p != nil {
		t.Error(p)
	}
}

func TestP1_S2_14(t *testing.T) {
	if p := P1(S2, 14); p == nil {
		t.Error(p)
	} else if *p != (Position{7, 3}) {
		t.Error(*p)
	}
}

func TestP1_X(t *testing.T) {
	if p := P1(X, 1000); p == nil {
		t.Error(p)
	} else if *p != (Position{57, 104}) {
		t.Error(*p)
	}
}

func TestP2_S3(t *testing.T) {
	if p := P2(S3, 10); p == nil {
		t.Error(p)
	} else if *p != (Position{6, 4}) {
		t.Error(*p)
	}
}

func TestP2_X(t *testing.T) {
	if p := P2(X, 100000); p == nil {
		t.Error(p)
	} else if *p != (Position{67, 74}) {
		t.Error(*p)
	}
}
