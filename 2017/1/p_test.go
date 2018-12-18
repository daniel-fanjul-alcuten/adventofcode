package p1

import "testing"

func TestP1_1122(t *testing.T) {
	if v := P1("1122"); v != 3 {
		t.Error(v)
	}
}

func TestP1_1111(t *testing.T) {
	if v := P1("1111"); v != 4 {
		t.Error(v)
	}
}

func TestP1_1234(t *testing.T) {
	if v := P1("1234"); v != 0 {
		t.Error(v)
	}
}

func TestP1_91212129(t *testing.T) {
	if v := P1("91212129"); v != 9 {
		t.Error(v)
	}
}

func TestP1_X(t *testing.T) {
	if v := P1(X); v != 1393 {
		t.Error(v)
	}
}

func TestP2_1212(t *testing.T) {
	if v := P2("1212"); v != 6 {
		t.Error(v)
	}
}

func TestP2_1221(t *testing.T) {
	if v := P2("1221"); v != 0 {
		t.Error(v)
	}
}

func TestP2_123425(t *testing.T) {
	if v := P2("123425"); v != 4 {
		t.Error(v)
	}
}

func TestP2_123123(t *testing.T) {
	if v := P2("123123"); v != 12 {
		t.Error(v)
	}
}

func TestP2_12131415(t *testing.T) {
	if v := P2("12131415"); v != 4 {
		t.Error(v)
	}
}

func TestP2_X(t *testing.T) {
	if v := P2(X); v != 1292 {
		t.Error(v)
	}
}
