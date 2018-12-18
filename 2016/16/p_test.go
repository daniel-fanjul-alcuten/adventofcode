package p

import (
	"testing"
)

func TestP1(t *testing.T) {
	if d, _ := P1(3, "1"); d != "100" {
		t.Error(d)
	}
	if d, _ := P1(3, "0"); d != "001" {
		t.Error(d)
	}
	if d, _ := P1(11, "11111"); d != "11111000000" {
		t.Error(d)
	}
	if d, _ := P1(25, "111100001010"); d != "1111000010100101011110000" {
		t.Error(d)
	}
	d, k := P1(12, "110010110100")
	if d != "110010110100" {
		t.Error(d)
	}
	if k != "100" {
		t.Error(k)
	}
	d, k = P1(20, "10000")
	if d != "10000011110010000111" {
		t.Error(d)
	}
	if k != "01100" {
		t.Error(k)
	}
	_, k = P1(272, "00101000101111010")
	if k != "10010100110011100" {
		t.Error(k)
	}
}

func TestP2(t *testing.T) {
	_, k := P1(35651584, "00101000101111010")
	if k != "01100100101101100" {
		t.Error(k)
	}
}
