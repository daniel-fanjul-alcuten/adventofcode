package p

import (
	"fmt"
	"testing"
)

func TestP1A_Samples(t *testing.T) {
	if s := fmt.Sprint(P1A("abcdef")); s != "false false" {
		t.Error(s)
	}
	if s := fmt.Sprint(P1A("bababc")); s != "true true" {
		t.Error(s)
	}
	if s := fmt.Sprint(P1A("abbcde")); s != "true false" {
		t.Error(s)
	}
	if s := fmt.Sprint(P1A("abcccd")); s != "false true" {
		t.Error(s)
	}
	if s := fmt.Sprint(P1A("aabcdd")); s != "true false" {
		t.Error(s)
	}
	if s := fmt.Sprint(P1A("abcdee")); s != "true false" {
		t.Error(s)
	}
	if s := fmt.Sprint(P1A("ababab")); s != "false true" {
		t.Error(s)
	}
}

func TestP1B_Samples(t *testing.T) {
	if k := P1B([]string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}); k != 12 {
		t.Error(k)
	}
}

func TestP1B_X(t *testing.T) {
	if k := P1B(X); k != 6972 {
		t.Error(k)
	}
}

func TestP2A_Samples(t *testing.T) {
	if s := fmt.Sprint(P2A("abcde", "axcye")); s != "false" {
		t.Error(s)
	}
	if s := fmt.Sprint(P2A("fghij", "fguij")); s != "fgijtrue" {
		t.Error(s)
	}
}

func TestP2B_X(t *testing.T) {
	if s := fmt.Sprint(P2B(X)); s != "[aixwcbzrmdvpsjfgllthdyoqe]" {
		t.Error(s)
	}
}
