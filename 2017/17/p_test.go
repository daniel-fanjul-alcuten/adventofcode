package p17

import (
	"testing"
)

func TestSample(t *testing.T) {
	s := New(3)
	for i := 0; i < 2017; i++ {
		s.Spin()
	}
	j := 1530
	for i, n := range s.Buffer {
		if n == 2017 {
			if i != j {
				t.Error(i)
			}
		}
	}
	if n := s.Buffer[j+1]; n != 638 {
		t.Error(n)
	}
}

func TestP1(t *testing.T) {
	s := New(376)
	for i := 0; i < 2017; i++ {
		s.Spin()
	}
	j := 735
	for i, n := range s.Buffer {
		if n == 2017 {
			if i != j {
				t.Error(i)
			}
		}
	}
	if n := s.Buffer[j+1]; n != 777 {
		t.Error(n)
	}
}

func TestP2(t *testing.T) {
	len := 1
	position := 0
	steps := 376
	positionOfZero := 0
	valueAfterZero := 0
	for i := 0; i < 50*1000*1000; i++ {
		p := (position + steps) % len
		if p == positionOfZero {
			valueAfterZero = len
		} else if p < positionOfZero {
			positionOfZero++
		}
		len++
		position = p + 1
	}
	if valueAfterZero != 39289581 {
		t.Error(valueAfterZero)
	}
}
