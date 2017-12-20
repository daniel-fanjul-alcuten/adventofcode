package p14

import (
	"testing"
)

func TestP1Sample(t *testing.T) {
	hh, u := Hash("flqrgnkx")
	if u != 8108 {
		t.Error(u)
	}
	m := ToMap(hh)
	if l := len(m); l != 8108 {
		t.Error(l)
	}
	if _, ok := m[Position{0, 0}]; !ok {
		t.Error(ok)
	}
	if _, ok := m[Position{1, 0}]; !ok {
		t.Error(ok)
	}
	if _, ok := m[Position{2, 0}]; ok {
		t.Error(ok)
	}
	if _, ok := m[Position{3, 0}]; !ok {
		t.Error(ok)
	}
	if _, ok := m[Position{4, 0}]; ok {
		t.Error(ok)
	}
	if _, ok := m[Position{5, 0}]; !ok {
		t.Error(ok)
	}
	if _, ok := m[Position{6, 0}]; ok {
		t.Error(ok)
	}
	if _, ok := m[Position{7, 0}]; ok {
		t.Error(ok)
	}
	if _, ok := m[Position{0, 1}]; ok {
		t.Error(ok)
	}
	if _, ok := m[Position{1, 1}]; !ok {
		t.Error(ok)
	}
	if _, ok := m[Position{2, 1}]; ok {
		t.Error(ok)
	}
	if _, ok := m[Position{3, 1}]; !ok {
		t.Error(ok)
	}
	if _, ok := m[Position{4, 1}]; ok {
		t.Error(ok)
	}
	if _, ok := m[Position{5, 1}]; !ok {
		t.Error(ok)
	}
	if _, ok := m[Position{6, 1}]; ok {
		t.Error(ok)
	}
	if _, ok := m[Position{7, 1}]; !ok {
		t.Error(ok)
	}
	if r := Regions(m); r != 1242 {
		t.Error(r)
	}
}

func TestP1(t *testing.T) {
	hh, u := Hash("uugsqrei")
	if u != 8194 {
		t.Error(u)
	}
	m := ToMap(hh)
	if l := len(m); l != 8194 {
		t.Error(l)
	}
	if r := Regions(m); r != 1141 {
		t.Error(r)
	}
}
