package kh

import (
	"testing"
)

func TestKnotHash(t *testing.T) {
	s := "199,0,255,136,174,254,227,16,51,85,1,2,22,17,7,192"
	if _, h := KnotHash([]byte(s)); h != "a9d0e68649d0174c8756a59ba21d4dc6" {
		t.Error(h)
	}
}
