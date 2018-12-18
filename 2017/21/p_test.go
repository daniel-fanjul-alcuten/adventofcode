package p21

import (
	"testing"
)

func TestFlip2x2(t *testing.T) {
	b := NewBuffer(2)
	p := Image{b, 0, 0, 2}
	p.Parse("12/34")
	if s := p.String(); s != "12/34" {
		t.Error(s)
	}
	p.Flip2x2()
	if s := p.String(); s != "21/43" {
		t.Error(s)
	}
	p.Flip2x2()
	if s := p.String(); s != "12/34" {
		t.Error(s)
	}
}

func TestRotate2x2(t *testing.T) {
	b := NewBuffer(2)
	p := Image{b, 0, 0, 2}
	p.Parse("12/34")
	if s := p.String(); s != "12/34" {
		t.Error(s)
	}
	p.Rotate2x2()
	if s := p.String(); s != "31/42" {
		t.Error(s)
	}
	p.Rotate2x2()
	if s := p.String(); s != "43/21" {
		t.Error(s)
	}
	p.Rotate2x2()
	if s := p.String(); s != "24/13" {
		t.Error(s)
	}
	p.Rotate2x2()
	if s := p.String(); s != "12/34" {
		t.Error(s)
	}
}

func TestFlip3x3(t *testing.T) {
	b := NewBuffer(3)
	p := Image{b, 0, 0, 3}
	p.Parse("123/456/789")
	if s := p.String(); s != "123/456/789" {
		t.Error(s)
	}
	p.Flip3x3()
	if s := p.String(); s != "321/654/987" {
		t.Error(s)
	}
	p.Flip3x3()
	if s := p.String(); s != "123/456/789" {
		t.Error(s)
	}
}

func TestRotate3x3(t *testing.T) {
	b := NewBuffer(3)
	p := Image{b, 0, 0, 3}
	p.Parse("123/456/789")
	if s := p.String(); s != "123/456/789" {
		t.Error(s)
	}
	p.Rotate3x3()
	if s := p.String(); s != "741/852/963" {
		t.Error(s)
	}
	p.Rotate3x3()
	if s := p.String(); s != "987/654/321" {
		t.Error(s)
	}
	p.Rotate3x3()
	if s := p.String(); s != "369/258/147" {
		t.Error(s)
	}
	p.Rotate3x3()
	if s := p.String(); s != "123/456/789" {
		t.Error(s)
	}
}

func TestS1(t *testing.T) {
	qq := Parse(S1)
	if l := len(qq); l != 12 {
		t.Error(l)
	}
	p := Image{NewBuffer(3), 0, 0, 3}
	p.Parse(".#./..#/###")
	if s := p.String(); s != ".#./..#/###" {
		t.Error(s)
	}
	p = p.Enhance(qq)
	if s := p.String(); s != "#..#/..../..../#..#" {
		t.Error(s)
	}
	p = p.Enhance(qq)
	if s := p.String(); s != "##.##./#..#../....../##.##./#..#../......" {
		t.Error(s)
	}
	if c := p.Count('#'); c != 12 {
		t.Error(c)
	}
}

func TestP1(t *testing.T) {
	qq := Parse(P1)
	if l := len(qq); l != 528 {
		t.Error(l)
	}
	p := Image{NewBuffer(3), 0, 0, 3}
	p.Parse(".#./..#/###")
	for i := 0; i < 5; i++ {
		p = p.Enhance(qq)
	}
	if c := p.Count('#'); c != 125 {
		t.Error(c)
	}
}

func TestP2(t *testing.T) {
	qq := Parse(P1)
	if l := len(qq); l != 528 {
		t.Error(l)
	}
	p := Image{NewBuffer(3), 0, 0, 3}
	p.Parse(".#./..#/###")
	for i := 0; i < 18; i++ {
		p = p.Enhance(qq)
	}
	if c := p.Count('#'); c != 1782917 {
		t.Error(c)
	}
}
