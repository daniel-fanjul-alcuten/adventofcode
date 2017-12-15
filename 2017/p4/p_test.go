package p4

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	if v := IsValid("aa bb cc dd ee"); !v {
		t.Error(v)
	}
	if v := IsValid("aa bb cc dd aa"); v {
		t.Error(v)
	}
	if v := IsValid("aa bb cc dd aaa"); !v {
		t.Error(v)
	}
}

func TestP1(t *testing.T) {
	if v := P1(x); v != 455 {
		t.Error(v)
	}
}

func TestIsValid2(t *testing.T) {
	if v := IsValid2("abcde fghij"); !v {
		t.Error(v)
	}
	if v := IsValid2("abcde xyz ecdab"); v {
		t.Error(v)
	}
	if v := IsValid2("a ab abc abd abf abj"); !v {
		t.Error(v)
	}
	if v := IsValid2("iiii oiii ooii oooi oooo"); !v {
		t.Error(v)
	}
	if v := IsValid2("oiii ioii iioi iiio"); v {
		t.Error(v)
	}
}

func TestP2(t *testing.T) {
	if v := P2(x); v != 186 {
		t.Error(v)
	}
}
