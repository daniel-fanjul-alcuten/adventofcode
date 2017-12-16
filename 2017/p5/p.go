package p5

type Jumps struct {
	Offset  int
	Offsets []int
}

func (j *Jumps) Jump() {
	f := j.Offset
	j.Offset += j.Offsets[f]
	j.Offsets[f] += 1
}

func (j *Jumps) Jump2() {
	f := j.Offset
	j.Offset += j.Offsets[f]
	if j.Offsets[f] >= 3 {
		j.Offsets[f] -= 1
	} else {
		j.Offsets[f] += 1
	}
}

func (j *Jumps) Done() bool {
	return j.Offset < 0 || j.Offset >= len(j.Offsets)
}

func P1(offsets []int) (n int) {
	j := Jumps{0, offsets}
	for !j.Done() {
		j.Jump()
		n++
	}
	return
}

func P2(offsets []int) (n int) {
	j := Jumps{0, offsets}
	for !j.Done() {
		j.Jump2()
		n++
	}
	return
}
