package p12

import (
	"testing"
)

func TestSample(t *testing.T) {
	pp := ParseN([]string{
		"0 <-> 2",
		"1 <-> 1",
		"2 <-> 0, 3, 4",
		"3 <-> 2, 4",
		"4 <-> 2, 3, 6",
		"5 <-> 6",
		"6 <-> 4, 5",
	})
	if s := pp["0"].Source; s != "0" {
		t.Error(s)
	}
	if _, ok := pp["0"].Targets["2"]; !ok {
		t.Error(ok)
	}
	if n := len(pp.Group("0")); n != 6 {
		t.Error(n)
	}
}

func TestP1(t *testing.T) {
	pp := ParseN(X)
	if n := len(pp.Group("0")); n != 378 {
		t.Error(n)
	}
}

func TestP2(t *testing.T) {
	pp := ParseN(X)
	n := 0
	for len(pp) > 0 {
		for s := range pp {
			for t := range pp.Group(s) {
				delete(pp, t)
			}
			n++
			break
		}
	}
	if n != 204 {
		t.Error(n)
	}
}
