package p3

type Direction int

const (
	StepToTheRight Direction = iota
	GoUp
	GoLeft
	GoDown
	GoRight
	NumDirections
)

type Position struct {
	X, Y int
}

type Record struct {
	Number                   int
	Position                 Position
	NextDirection            Direction
	PendingSteps             int
	PendingStepsPerDirection int
}

func FirstRecord() Record {
	return Record{1, Position{0, 0}, StepToTheRight, 1, 0}
}

func (r Record) Next() Record {

	r.Number += 1
	switch r.NextDirection {
	case StepToTheRight:
		r.Position.X += 1
	case GoUp:
		r.Position.Y -= 1
	case GoLeft:
		r.Position.X -= 1
	case GoDown:
		r.Position.Y += 1
	case GoRight:
		r.Position.X += 1
	}
	r.PendingSteps -= 1

	if r.PendingSteps <= 0 {
		if r.NextDirection == StepToTheRight {
			r.PendingStepsPerDirection += 2
		}
		r.NextDirection = (r.NextDirection + 1) % NumDirections
		if r.NextDirection == StepToTheRight {
			r.PendingSteps = 1
		} else if r.NextDirection == GoUp {
			r.PendingSteps = r.PendingStepsPerDirection - 1
		} else {
			r.PendingSteps = r.PendingStepsPerDirection
		}
	}
	return r
}

func FirstRecord2(prev map[Position]int) Record {
	r := FirstRecord()
	prev[r.Position] = r.Number
	return r
}

func (r Record) Next2(prev map[Position]int) Record {
	r = r.Next()
	r.Number = 0
	for _, x := range []int{-1, 0, +1} {
		for _, y := range []int{-1, 0, +1} {
			r.Number += prev[Position{r.Position.X + x, r.Position.Y + y}]
		}
	}
	prev[r.Position] = r.Number
	return r
}

func P1(i int) (s int) {
	r := FirstRecord()
	for r.Number < i {
		r = r.Next()
	}
	if r.Position.X > 0 {
		s += r.Position.X
	}
	if r.Position.X < 0 {
		s -= r.Position.X
	}
	if r.Position.Y > 0 {
		s += r.Position.Y
	}
	if r.Position.Y < 0 {
		s -= r.Position.Y
	}
	return
}

func P2(max int) int {
	prev := map[Position]int{}
	r := FirstRecord()
	prev[r.Position] = r.Number
	for r.Number < max {
		r = r.Next2(prev)
		prev[r.Position] = r.Number
	}
	return r.Number
}
