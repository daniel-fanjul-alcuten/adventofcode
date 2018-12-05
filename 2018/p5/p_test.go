package p

import (
	"testing"
)

func TestP1_S(t *testing.T) {
	if s := P1("aA"); s != "" {
		t.Error(s)
	}
	if s := P1("abBA"); s != "" {
		t.Error(s)
	}
	if s := P1("abAB"); s != "abAB" {
		t.Error(s)
	}
	if s := P1("aabAAB"); s != "aabAAB" {
		t.Error(s)
	}
	if s := P1("dabAcCaCBAcCcaDA"); s != "dabCBAcaDA" {
		t.Error(s)
	}
}

func TestP1_X(t *testing.T) {
	if n := len(P1(X)); n != 9808 {
		t.Error(n)
	}
}

func TestP2_S(t *testing.T) {
	if s := P2("dabAcCaCBAcCcaDA"); s != "daDA" {
		t.Error(s)
	}
}

func TestP2_X(t *testing.T) {
	if n := len(P2(X)); n != 6484 {
		t.Error(n)
	}
}
