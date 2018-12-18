package p20

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestS1(t *testing.T) {
	tt := ParticlesByCloseToZero(ParseN(S1))
	sort.Sort(&tt)
	if id := tt[0].ID; id != 0 {
		t.Error(id)
	}
}

func TestP1(t *testing.T) {
	tt := ParticlesByCloseToZero(ParseN(P1))
	sort.Sort(&tt)
	if id := tt[0].ID; id != 243 {
		t.Error(id)
	}
}

func TestCollisionTimes(t *testing.T) {
	t1, t2 := CollisionTimes(-6, 3, 0, -4, 2, 0)
	if t1 != 2 {
		t.Error(t1)
	}
	if t2 != 2 {
		t.Error(t2)
	}
	t1, t2 = CollisionTimes(-6, 3, 0, 3, -1, 0)
	if t1 != 2.25 {
		t.Error(t1)
	}
	if t2 != 2.25 {
		t.Error(t2)
	}
	t1, t2 = CollisionTimes(-6, 3, 0, -4, 2, 1)
	if !math.IsNaN(t1) {
		t.Error(t1)
	}
	if !math.IsNaN(t2) {
		t.Error(t2)
	}
	t1, t2 = CollisionTimes(-7, 0, 0, 4, 0, 0)
	if !math.IsInf(t1, 1) {
		t.Error(t1)
	}
	if !math.IsInf(t2, 1) {
		t.Error(t2)
	}
	t1, t2 = CollisionTimes(4, 0, 0, 4, 0, 0)
	if t1 != 0 {
		t.Error(t1)
	}
	if t2 != 0 {
		t.Error(t2)
	}
}

func TestCollisionSet(t *testing.T) {
	if m := fmt.Sprint(CollisionSet(-6, 3, 0, -4, 2, 0)); m != "map[2:{}]" {
		t.Error(m)
	}
	if m := fmt.Sprint(CollisionSet(-6, 3, 0, 3, -1, 0)); m != "map[]" {
		t.Error(m)
	}
	if m := fmt.Sprint(CollisionSet(-6, 3, 0, -4, 2, 1)); m != "map[]" {
		t.Error(m)
	}
	if m := fmt.Sprint(CollisionSet(-7, 0, 0, 4, 0, 0)); m != "map[]" {
		t.Error(m)
	}
	if m := fmt.Sprint(CollisionSet(4, 0, 0, 4, 0, 0)); m != "map[0:{}]" {
		t.Error(m)
	}
	if m := fmt.Sprint(CollisionSet(-6, 3, 0, -6, 2, 0)); m != "map[0:{}]" {
		t.Error(m)
	}
	if m := fmt.Sprint(CollisionSet(-6, 3, 1, -6, 2, 2)); m != "map[0:{} 1:{}]" && m != "map[1:{} 0:{}]" {
		t.Error(m)
	}
	if m := fmt.Sprint(CollisionSet(-6, 3, 1, -7, 2, 2)); m != "map[2:{}]" {
		t.Error(m)
	}
	if m := fmt.Sprint(CollisionSet(-6, 3, 1, -8, 2, 2)); m != "map[]" {
		t.Error(m)
	}
}

func TestP2(t *testing.T) {
	tt := ParseN(P1)
	i := 0
	for i < len(tt) {
		p, j, ok := tt[i], i+1, true
		for j < len(tt) {
			q := tt[j]
			if p.Collides(q) {
				ok, tt = false, append(tt[:j], tt[j+1:]...)
			} else {
				j++
			}
		}
		if ok {
			i++
		} else {
			tt = append(tt[:i], tt[i+1:]...)
		}
	}
	if l := len(tt); l != 648 {
		t.Error(l)
	}
}
