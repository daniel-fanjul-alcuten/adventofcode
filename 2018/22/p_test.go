package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	risk, mins := P12(P{}, P{10, 10}, 510)
	if risk != 114 {
		t.Error(risk)
	}
	if mins != 45 {
		t.Error(mins)
	}
}

func TestP1_X(t *testing.T) {
	risk, mins := P12(P{}, P{9, 751}, 11817)
	if risk != 7402 {
		t.Error(risk)
	}
	if mins != 1025 {
		t.Error(mins)
	}
}
