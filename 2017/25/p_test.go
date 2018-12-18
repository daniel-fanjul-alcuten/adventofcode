package p25

import (
	"testing"
)

func TestS1(t *testing.T) {
	p := Tape{}
	p.Execute(S1)
	if k := p.Checksum(); k != 3 {
		t.Error(k)
	}
}

func TestP1(t *testing.T) {
	p := Tape{}
	p.Execute(P1)
	if k := p.Checksum(); k != 4230 {
		t.Error(k)
	}
}
