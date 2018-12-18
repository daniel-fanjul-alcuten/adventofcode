package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if s := P1(S); s != "CABDFE" {
		t.Error(s)
	}
}

func TestP1_X(t *testing.T) {
	if s := P1(X); s != "OCPUEFIXHRGWDZABTQJYMNKVSL" {
		t.Error(s)
	}
}

func TestP2_S(t *testing.T) {
	if s, dur := P2(S, 2, 0); s != "CABFDE" {
		t.Error(s)
	} else if dur != 15 {
		t.Error(dur)
	}
}

func TestP2_X(t *testing.T) {
	if s, dur := P2(X, 5, 60); s != "OPCUXEHFIRWZGDABTQYJMNKVSL" {
		t.Error(s)
	} else if dur != 991 {
		t.Error(dur)
	}
}
