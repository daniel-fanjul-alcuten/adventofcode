package p

import "testing"

func TestP1_S(t *testing.T) {
	for i := 1; i <= 10; i++ {
		if n := P1(i, false); n != 1 {
			t.Error(i, n)
		}
	}
	for i := 11; i <= 30; i++ {
		if n := P1(i, false); n != 2 {
			t.Error(i, n)
		}
	}
	for i := 31; i <= 40; i++ {
		if n := P1(i, false); n != 3 {
			t.Error(i, n)
		}
	}
	for i := 41; i <= 70; i++ {
		if n := P1(i, false); n != 4 {
			t.Error(i, n)
		}
	}
	for i := 71; i <= 120; i++ {
		if n := P1(i, false); n != 6 {
			t.Error(i, n)
		}
	}
	for i := 121; i <= 150; i++ {
		if n := P1(i, false); n != 8 {
			t.Error(i, n)
		}
	}
}

func TestP1_X(t *testing.T) {
	t.Parallel()
	if n := P1(36000000, false); n != 831600 {
		t.Error(n)
	}
}

func TestP2_X(t *testing.T) {
	t.Parallel()
	if n := P1(36000000, true); n != 884520 {
		t.Error(n)
	}
}
