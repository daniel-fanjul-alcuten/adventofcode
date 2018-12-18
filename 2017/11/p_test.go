package p11

import (
	"testing"
)

func TestSample(t *testing.T) {
	if d := (Position{}).AddN("ne,ne,ne").Distance(); d != 3 {
		t.Error(d)
	}
	if d := (Position{}).AddN("ne,ne,sw,sw").Distance(); d != 0 {
		t.Error(d)
	}
	if d := (Position{}).AddN("ne,ne,s,s").Distance(); d != 2 {
		t.Error(d)
	}
	if d := (Position{}).AddN("se,sw,se,sw,sw").Distance(); d != 3 {
		t.Error(d)
	}
}

func TestP1(t *testing.T) {
	if d := (Position{}).AddN(X).Distance(); d != 761 {
		t.Error(d)
	}
	if d := (Position{}).AddN(X).D; d != 1542 {
		t.Error(d)
	}
}
