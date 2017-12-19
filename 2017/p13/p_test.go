package p13

import (
	"math"
	"testing"
)

func TestSample1(t *testing.T) {
	ll := New(map[int]int{
		0: 3,
		1: 2,
		4: 4,
		6: 4,
	})
	d := ll.MaxDepth()
	if d != 6 {
		t.Error(d)
	}
	ll.AdvancePosition()
	ll.AdvanceScanners()
	if v := ll.Picosecond; v != 0 {
		t.Error(v)
	}
	if v := ll.Position; v != 0 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[0].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[1].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[4].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[6].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.Severity; v != 0 {
		t.Error(v)
	}
	ll.AdvancePosition()
	ll.AdvanceScanners()
	if v := ll.Picosecond; v != 1 {
		t.Error(v)
	}
	if v := ll.Position; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[0].Scanner; v != 2 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[1].Scanner; v != 0 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[4].Scanner; v != 2 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[6].Scanner; v != 2 {
		t.Error(v)
	}
	if v := ll.Severity; v != 0 {
		t.Error(v)
	}
	ll.AdvancePosition()
	ll.AdvanceScanners()
	if v := ll.Picosecond; v != 2 {
		t.Error(v)
	}
	if v := ll.Position; v != 2 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[0].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[1].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[4].Scanner; v != 3 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[6].Scanner; v != 3 {
		t.Error(v)
	}
	if v := ll.Severity; v != 0 {
		t.Error(v)
	}
	ll.AdvancePosition()
	ll.AdvanceScanners()
	if v := ll.Picosecond; v != 3 {
		t.Error(v)
	}
	if v := ll.Position; v != 3 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[0].Scanner; v != 0 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[1].Scanner; v != 0 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[4].Scanner; v != 2 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[6].Scanner; v != 2 {
		t.Error(v)
	}
	if v := ll.Severity; v != 0 {
		t.Error(v)
	}
	ll.AdvancePosition()
	ll.AdvanceScanners()
	if v := ll.Picosecond; v != 4 {
		t.Error(v)
	}
	if v := ll.Position; v != 4 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[0].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[1].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[4].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[6].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.Severity; v != 0 {
		t.Error(v)
	}
	ll.AdvancePosition()
	ll.AdvanceScanners()
	if v := ll.Picosecond; v != 5 {
		t.Error(v)
	}
	if v := ll.Position; v != 5 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[0].Scanner; v != 2 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[1].Scanner; v != 0 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[4].Scanner; v != 0 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[6].Scanner; v != 0 {
		t.Error(v)
	}
	if v := ll.Severity; v != 0 {
		t.Error(v)
	}
	ll.AdvancePosition()
	ll.AdvanceScanners()
	if v := ll.Picosecond; v != 6 {
		t.Error(v)
	}
	if v := ll.Position; v != 6 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[0].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[1].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[4].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[6].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.Severity; v != 24 {
		t.Error(v)
	}
}

func TestSample2(t *testing.T) {
	ll := New(map[int]int{
		0: 3,
		1: 2,
		4: 4,
		6: 4,
	})
	d := ll.MaxDepth()
	if d != 6 {
		t.Error(d)
	}
	for _ = range [10]int{} {
		ll.AdvanceScanners()
	}
	for _ = range [7]int{} {
		ll.AdvancePosition()
		ll.AdvanceScanners()
	}
	if v := ll.Picosecond; v != 16 {
		t.Error(v)
	}
	if v := ll.Position; v != 6 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[0].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[1].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[4].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.LayersByDepth[6].Scanner; v != 1 {
		t.Error(v)
	}
	if v := ll.Severity; v != 0 {
		t.Error(v)
	}
}

func TestP1(t *testing.T) {
	ll := New(X)
	d := ll.MaxDepth()
	if d != 96 {
		t.Error(d)
	}
	for {
		ll.AdvancePosition()
		ll.AdvanceScanners()
		if ll.Position == d {
			break
		}
	}
	if v := ll.Severity; v != 1612 {
		t.Error(v)
	}
}

func TestP2(t *testing.T) {
	ll := New(X)
	d := ll.MaxDepth()
	for delay := 0; delay < math.MaxInt32-1; delay++ {
		ll2 := ll.Clone()
		for {
			ll.AdvancePosition()
			if ll.Caught {
				break
			}
			if ll.Position >= d {
				if delay != 3907994 {
					t.Error(delay)
				}
				return
			}
			ll.AdvanceScanners()
		}
		ll = ll2
		ll.AdvanceScanners()
	}
	t.Error("Failed")
}
