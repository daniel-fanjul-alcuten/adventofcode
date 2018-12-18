package p

import "testing"

func TestP1_S1(t *testing.T) {
	t.Parallel()
	if n := P1("abcdef"); n != 609043 {
		t.Error(n)
	}
}

func TestP1_S2(t *testing.T) {
	t.Parallel()
	if n := P1("pqrstuv"); n != 1048970 {
		t.Error(n)
	}
}

func TestP1_X(t *testing.T) {
	t.Parallel()
	if n := P1("yzbqklnj"); n != 282749 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	t.Parallel()
	if n := P2("yzbqklnj"); n != 9962624 {
		t.Error(n)
	}
}
