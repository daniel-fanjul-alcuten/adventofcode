package p8

import (
	"math"
	"testing"
)

func TestParse1(t *testing.T) {
	k := Parse1(S[0])
	if v := k.Variable; v != "b" {
		t.Error(v)
	}
	if v := k.Operation; v != "inc" {
		t.Error(v)
	}
	if v := k.Value; v != 5 {
		t.Error(v)
	}
	if v := k.CondVariable; v != "a" {
		t.Error(v)
	}
	if v := k.CondOperator; v != ">" {
		t.Error(v)
	}
	if v := k.CondValue; v != 1 {
		t.Error(v)
	}
	k = Parse1(S[2])
	if v := k.Variable; v != "c" {
		t.Error(v)
	}
	if v := k.Operation; v != "dec" {
		t.Error(v)
	}
	if v := k.Value; v != -10 {
		t.Error(v)
	}
	if v := k.CondVariable; v != "a" {
		t.Error(v)
	}
	if v := k.CondOperator; v != ">=" {
		t.Error(v)
	}
	if v := k.CondValue; v != 1 {
		t.Error(v)
	}
}

func TestSample(t *testing.T) {
	rr := Records{}
	rr.ApplyN(ParseN(S))
	if v := rr["a"]; v != 1 {
		t.Error(v)
	}
	if v := rr["b"]; v != 0 {
		t.Error(v)
	}
	if v := rr["c"]; v != -10 {
		t.Error(v)
	}
}

func TestP1(t *testing.T) {
	rr := Records{}
	rr.ApplyN(ParseN(X))
	if v := rr.Max(); v != 5075 {
		t.Error(v)
	}
}

func TestP2(t *testing.T) {
	rr := Records{}
	max := math.MinInt32
	for _, k := range ParseN(X) {
		rr.Apply1(k)
		if m := rr.Max(); m > max {
			max = m
		}
	}
	if v := max; v != 7310 {
		t.Error(v)
	}
}
