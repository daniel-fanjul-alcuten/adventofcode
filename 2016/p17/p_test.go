package p

import (
	"testing"
)

func TestP1_S1(t *testing.T) {
	if p := P1(P{0, 0}, P{3, 3}, "hijkl"); p != "" {
		t.Error(p)
	}
}

func TestP1_S2(t *testing.T) {
	if p := P1(P{0, 0}, P{3, 3}, "ihgpwlah"); p != "DDRRRD" {
		t.Error(p)
	}
}

func TestP1_S3(t *testing.T) {
	if p := P1(P{0, 0}, P{3, 3}, "kglvqrro"); p != "DDUDRLRRUDRD" {
		t.Error(p)
	}
}

func TestP1_S4(t *testing.T) {
	if p := P1(P{0, 0}, P{3, 3}, "ulqzkmiv"); p != "DRURDRUDDLLDLUURRDULRLDUUDDDRR" {
		t.Error(p)
	}
}

func TestP1_X(t *testing.T) {
	if p := P1(P{0, 0}, P{3, 3}, "vwbaicqe"); p != "DRDRULRDRD" {
		t.Error(p)
	}
}

func TestP2_S1(t *testing.T) {
	if p := len(P2(P{0, 0}, P{3, 3}, "ihgpwlah")); p != 370 {
		t.Error(p)
	}
}

func TestP2_S2(t *testing.T) {
	if p := len(P2(P{0, 0}, P{3, 3}, "kglvqrro")); p != 492 {
		t.Error(p)
	}
}

func TestP2_S3(t *testing.T) {
	if p := len(P2(P{0, 0}, P{3, 3}, "ulqzkmiv")); p != 830 {
		t.Error(p)
	}
}

func TestP2_X(t *testing.T) {
	if p := len(P2(P{0, 0}, P{3, 3}, "vwbaicqe")); p != 384 {
		t.Error(p)
	}
}
