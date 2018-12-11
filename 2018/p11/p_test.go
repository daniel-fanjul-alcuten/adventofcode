package p

import (
	"testing"
)

func TestPowerLevel(t *testing.T) {
	t.Parallel()
	if r := PowerLevel(3, 5, 8); r != 4 {
		t.Error(r)
	}
	if r := PowerLevel(122, 79, 57); r != -5 {
		t.Error(r)
	}
	if r := PowerLevel(217, 196, 39); r != 0 {
		t.Error(r)
	}
	if r := PowerLevel(101, 153, 71); r != 4 {
		t.Error(r)
	}
}

func TestP1_S1(t *testing.T) {
	t.Parallel()
	x, y, pl := P1(300, 300, 18)
	if x != 33 {
		t.Error(x)
	}
	if y != 45 {
		t.Error(y)
	}
	if pl != 29 {
		t.Error(pl)
	}
}

func TestP1_S2(t *testing.T) {
	t.Parallel()
	x, y, pl := P1(300, 300, 42)
	if x != 21 {
		t.Error(x)
	}
	if y != 61 {
		t.Error(y)
	}
	if pl != 30 {
		t.Error(pl)
	}
}

func TestP1_X(t *testing.T) {
	t.Parallel()
	x, y, pl := P1(300, 300, 4172)
	if x != 243 {
		t.Error(x)
	}
	if y != 43 {
		t.Error(y)
	}
	if pl != 29 {
		t.Error(pl)
	}
}

func TestP2_S1(t *testing.T) {
	t.Parallel()
	x, y, z, pl := P2(300, 18)
	if x != 90 {
		t.Error(x)
	}
	if y != 269 {
		t.Error(y)
	}
	if z != 16 {
		t.Error(z)
	}
	if pl != 113 {
		t.Error(pl)
	}
}

func TestP2_S2(t *testing.T) {
	t.Parallel()
	x, y, z, pl := P2(300, 42)
	if x != 232 {
		t.Error(x)
	}
	if y != 251 {
		t.Error(y)
	}
	if z != 12 {
		t.Error(z)
	}
	if pl != 119 {
		t.Error(pl)
	}
}

func TestP2_X(t *testing.T) {
	t.Parallel()
	x, y, z, pl := P2(300, 4172)
	if x != 236 {
		t.Error(x)
	}
	if y != 151 {
		t.Error(y)
	}
	if z != 15 {
		t.Error(z)
	}
	if pl != 127 {
		t.Error(pl)
	}
}
