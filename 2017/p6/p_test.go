package p6

import (
	"testing"
)

func TestExample(t *testing.T) {
	k := Banks{0, 2, 7, 0}
	if v := k.MaxIndex(); v != 2 {
		t.Error(v)
	}
	k.Redistribute(2)
	if s := k.String(); s != "2,4,1,2" {
		t.Error(s)
	}
	if v := k.MaxIndex(); v != 1 {
		t.Error(v)
	}
	k.Redistribute(1)
	if s := k.String(); s != "3,1,2,3" {
		t.Error(s)
	}
	if v := k.MaxIndex(); v != 0 {
		t.Error(v)
	}
	k.Redistribute(0)
	if s := k.String(); s != "0,2,3,4" {
		t.Error(s)
	}
	if v := k.MaxIndex(); v != 3 {
		t.Error(v)
	}
	k.Redistribute(3)
	if s := k.String(); s != "1,3,4,1" {
		t.Error(s)
	}
	if v := k.MaxIndex(); v != 2 {
		t.Error(v)
	}
	k.Redistribute(2)
	if s := k.String(); s != "2,4,1,2" {
		t.Error(s)
	}
}

func TestP1(t *testing.T) {
	k := Banks(X())
	if v := P1(k); v != 5042 {
		t.Error(v)
	}
}

func TestP2(t *testing.T) {
	k := Banks(X())
	if v := P2(k); v != 1086 {
		t.Error(v)
	}
}
