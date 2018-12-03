package p

import (
	"fmt"
	"testing"
)

func TestClaimParse(t *testing.T) {
	var c Claim
	if err := c.Parse("123 @ 3,2: 5x4"); err != nil {
		t.Error(err)
	}
	if s := fmt.Sprintf("%#v", c); s != `p.Claim{ID:"123", Left:3, Top:2, Width:5, Height:4}` {
		t.Error(s)
	}
}

func TestP1_Sample(t *testing.T) {
	var c1 Claim
	if err := c1.Parse("#1 @ 1,3: 4x4"); err != nil {
		t.Error(err)
	}
	var c2 Claim
	if err := c2.Parse("#2 @ 3,1: 4x4"); err != nil {
		t.Error(err)
	}
	var c3 Claim
	if err := c3.Parse("#3 @ 5,5: 2x2"); err != nil {
		t.Error(err)
	}
	if n := P1(c1, c2, c3); n != 4 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	cc := make([]Claim, len(X))
	for i, s := range X {
		if err := cc[i].Parse(s); err != nil {
			t.Error(err)
		}
	}
	if n := P1(cc...); n != 121259 {
		t.Error(n)
	}
}

func TestP2_Samples(t *testing.T) {
	var c1 Claim
	if err := c1.Parse("#1 @ 1,3: 4x4"); err != nil {
		t.Error(err)
	}
	var c2 Claim
	if err := c2.Parse("#2 @ 3,1: 4x4"); err != nil {
		t.Error(err)
	}
	var c3 Claim
	if err := c3.Parse("#3 @ 5,5: 2x2"); err != nil {
		t.Error(err)
	}
	if n := P1(c1, c2, c3); n != 4 {
		t.Error(n)
	}
	if s := fmt.Sprint(P2(c1, c2, c3)); s != "[#3]" {
		t.Error(s)
	}
}

func TestP2_X(t *testing.T) {
	cc := make([]Claim, len(X))
	for i, s := range X {
		if err := cc[i].Parse(s); err != nil {
			t.Error(err)
		}
	}
	if s := fmt.Sprint(P2(cc...)); s != "[#239]" {
		t.Error(s)
	}
}
