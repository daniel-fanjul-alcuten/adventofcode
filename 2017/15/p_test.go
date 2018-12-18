package p15

import (
	"testing"
)

func TestNext(t *testing.T) {
	a, b := Generator{65, 16807, 4}, Generator{8921, 48271, 8}
	if n := a.Next(); n != 1092455 {
		t.Error(n)
	}
	if n := a.Next(); n != 1181022009 {
		t.Error(n)
	}
	if n := a.Next(); n != 245556042 {
		t.Error(n)
	}
	if n := a.Next(); n != 1744312007 {
		t.Error(n)
	}
	if n := a.Next(); n != 1352636452 {
		t.Error(n)
	}
	if n := b.Next(); n != 430625591 {
		t.Error(n)
	}
	if n := b.Next(); n != 1233683848 {
		t.Error(n)
	}
	if n := b.Next(); n != 1431495498 {
		t.Error(n)
	}
	if n := b.Next(); n != 137874439 {
		t.Error(n)
	}
	if n := b.Next(); n != 285222916 {
		t.Error(n)
	}
}

func TestNext2(t *testing.T) {
	a, b := Generator{65, 16807, 4}, Generator{8921, 48271, 8}
	if n := a.Next2(); n != 1352636452 {
		t.Error(n)
	}
	if n := a.Next2(); n != 1992081072 {
		t.Error(n)
	}
	if n := a.Next2(); n != 530830436 {
		t.Error(n)
	}
	if n := a.Next2(); n != 1980017072 {
		t.Error(n)
	}
	if n := a.Next2(); n != 740335192 {
		t.Error(n)
	}
	if n := b.Next2(); n != 1233683848 {
		t.Error(n)
	}
	if n := b.Next2(); n != 862516352 {
		t.Error(n)
	}
	if n := b.Next2(); n != 1159784568 {
		t.Error(n)
	}
	if n := b.Next2(); n != 1616057672 {
		t.Error(n)
	}
	if n := b.Next2(); n != 412269392 {
		t.Error(n)
	}
}

func TestSampleP1(t *testing.T) {
	a, b, c := Generator{65, 16807, 4}, Generator{8921, 48271, 8}, 0
	for i := 0; i < 40*1000*1000; i++ {
		n, m := a.Next(), b.Next()
		c += Compare(n, m)
	}
	if c != 588 {
		t.Error(c)
	}
}

func TestSampleP2_A(t *testing.T) {
	a, b, c := Generator{65, 16807, 4}, Generator{8921, 48271, 8}, 0
	for i := 0; i < 1055; i++ {
		n, m := a.Next2(), b.Next2()
		c += Compare(n, m)
	}
	if c != 0 {
		t.Error(c)
	}
	n, m := a.Next2(), b.Next2()
	c += Compare(n, m)
	if c != 1 {
		t.Error(c)
	}
}

func TestSampleP2_B(t *testing.T) {
	a, b, c := Generator{65, 16807, 4}, Generator{8921, 48271, 8}, 0
	for i := 0; i < 5*1000*1000; i++ {
		n, m := a.Next2(), b.Next2()
		c += Compare(n, m)
	}
	if c != 309 {
		t.Error(c)
	}
}

func TestP1(t *testing.T) {
	a, b, c := Generator{873, 16807, 4}, Generator{583, 48271, 8}, 0
	for i := 0; i < 40*1000*1000; i++ {
		n, m := a.Next(), b.Next()
		c += Compare(n, m)
	}
	if c != 631 {
		t.Error(c)
	}
}

func TestP2(t *testing.T) {
	a, b, c := Generator{873, 16807, 4}, Generator{583, 48271, 8}, 0
	for i := 0; i < 5*1000*1000; i++ {
		n, m := a.Next2(), b.Next2()
		c += Compare(n, m)
	}
	if c != 279 {
		t.Error(c)
	}
}
