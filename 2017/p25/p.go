package p25

type Tape map[int]byte

func (t Tape) Execute(m Machine) {
	s, p := m.State, 0
	for i := 0; i < m.Steps; i++ {
		op := m.States[s][t[p]]
		t[p], p, s = op.Write, p+op.Move, op.Next
	}
}

func (t Tape) Checksum() (k int) {
	for _, b := range t {
		if b == 1 {
			k++
		}
	}
	return
}
