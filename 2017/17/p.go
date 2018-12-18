package p17

type Spinlock struct {
	Buffer   []int
	Position int
	Steps    int
}

func New(steps int) Spinlock {
	return Spinlock{[]int{0}, 0, steps}
}

func (s *Spinlock) Spin() {
	p := (s.Position + s.Steps) % len(s.Buffer)
	b := make([]int, 0, len(s.Buffer)+1)
	b = append(b, s.Buffer[:p+1]...)
	b = append(b, len(s.Buffer))
	b = append(b, s.Buffer[p+1:]...)
	s.Buffer = b
	s.Position = p + 1
}
