package p5

import (
	"testing"
)

func TestJump(t *testing.T) {
	j := Jumps{0, []int{0, 3, 0, 1, -3}}
	j.Jump()
	if v := j.Offset; v != 0 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump()
	if v := j.Offset; v != 1 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump()
	if v := j.Offset; v != 4 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump()
	if v := j.Offset; v != 1 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump()
	if v := j.Offset; v != 5 {
		t.Error(v)
	}
	if v := j.Done(); !v {
		t.Error(v)
	}
}

func TestP1(t *testing.T) {
	if v := P1(X()); v != 343364 {
		t.Error(v)
	}
}

func TestJump2(t *testing.T) {
	j := Jumps{0, []int{0, 3, 0, 1, -3}}
	j.Jump2()
	if v := j.Offset; v != 0 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump2()
	if v := j.Offset; v != 1 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump2()
	if v := j.Offset; v != 4 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump2()
	if v := j.Offset; v != 1 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump2()
	if v := j.Offset; v != 3 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump2()
	if v := j.Offset; v != 4 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump2()
	if v := j.Offset; v != 2 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump2()
	if v := j.Offset; v != 2 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump2()
	if v := j.Offset; v != 3 {
		t.Error(v)
	}
	if v := j.Done(); v {
		t.Error(v)
	}
	j.Jump2()
	if v := j.Offset; v != 5 {
		t.Error(v)
	}
	if v := j.Done(); !v {
		t.Error(v)
	}
	if v := j.Offsets[0]; v != 2 {
		t.Error(v)
	}
	if v := j.Offsets[1]; v != 3 {
		t.Error(v)
	}
	if v := j.Offsets[2]; v != 2 {
		t.Error(v)
	}
	if v := j.Offsets[3]; v != 3 {
		t.Error(v)
	}
	if v := j.Offsets[4]; v != -1 {
		t.Error(v)
	}
}

func TestP2(t *testing.T) {
	if v := P2(X()); v != 25071947 {
		t.Error(v)
	}
}
