package p7

import (
	"testing"
)

func TestS1(t *testing.T) {
	Parse(S1)
}

func TestP1(t *testing.T) {
	defer func() {
		n := recover()
		if n != "areod" {
			t.Error(n)
		}
	}()
	Parse(P1)
}
