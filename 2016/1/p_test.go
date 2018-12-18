package p1

import "testing"

const P1 = "L5, R1, R3, L4, R3, R1, L3, L2, R3, L5, L1, L2, R5, L1, R5, R1, L4, R1, R3, L4, L1, R2, R5, R3, R1, R1, L1, R1, L1, L2, L1, R2, L5, L188, L4, R1, R4, L3, R47, R1, L1, R77, R5, L2, R1, L2, R4, L5, L1, R3, R187, L4, L3, L3, R2, L3, L5, L4, L4, R1, R5, L4, L3, L3, L3, L2, L5, R1, L2, R5, L3, L4, R4, L5, R3, R4, L2, L1, L4, R1, L3, R1, R3, L2, R1, R4, R5, L3, R5, R3, L3, R4, L2, L5, L1, L1, R3, R1, L4, R3, R3, L2, R5, R4, R1, R3, L4, R3, R3, L2, L4, L5, R1, L4, L5, R4, L2, L1, L3, L3, L5, R3, L4, L3, R5, R4, R2, L4, R2, R3, L3, R4, L1, L3, R2, R1, R5, L4, L5, L5, R4, L5, L2, L4, R4, R4, R1, L3, L2, L4, R3"

func TestS1(t *testing.T) {
	p := New()
	p.Exec1N("R2, L3")
	if d := p.Distance(); d != 5 {
		t.Error(d)
	}
}

func TestS2(t *testing.T) {
	p := New()
	p.Exec1N("R2, R2, R2")
	if d := p.Distance(); d != 2 {
		t.Error(d)
	}
}

func TestS3(t *testing.T) {
	p := New()
	p.Exec1N("R5, L5, R5, R3")
	if d := p.Distance(); d != 12 {
		t.Error(d)
	}
}

func TestP1(t *testing.T) {
	p := New()
	p.Exec1N(P1)
	if d := p.Distance(); d != 273 {
		t.Error(d)
	}
}

func TestP2(t *testing.T) {
	p, m := New(), map[P]struct{}{}
	p.Exec2N(P1, m)
	if d := p.Distance(); d != 115 {
		t.Error(d)
	}
}
