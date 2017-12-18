package p10

import "encoding/hex"

type Sequence struct {
	Data     []byte
	Position int
	Skip     int
}

func New(n int) (s Sequence) {
	s.Data = make([]byte, n)
	for i := range s.Data {
		s.Data[i] = byte(i)
	}
	return
}

func (s *Sequence) Apply1(l int) {
	i, j := s.Position, (s.Position+l-1+len(s.Data))%len(s.Data)
	for t := l; t > 1; t -= 2 {
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
		i, j = (i+1)%len(s.Data), (j-1+len(s.Data))%len(s.Data)
	}
	s.Position = (s.Position + l + s.Skip) % len(s.Data)
	s.Skip += 1
}

func (s *Sequence) ApplyN(ll []int) {
	for _, l := range ll {
		s.Apply1(l)
	}
}

func (s *Sequence) Hash1() int {
	return int(s.Data[0]) * int(s.Data[1])
}

func (s *Sequence) ApplyN2(ll []byte) {
	for _ = range [64]int{} {
		for _, l := range append(ll, 17, 31, 73, 47, 23) {
			s.Apply1(int(l))
		}
	}
}

func (s *Sequence) Hash2() string {
	h := []byte{}
	for i := 0; i < len(s.Data); i += 16 {
		t := s.Data[i]
		for j := i + 1; j < i+16; j++ {
			t ^= s.Data[j]
		}
		h = append(h, t)
	}
	return hex.EncodeToString(h)
}
