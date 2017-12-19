package p13

import "math"

type Layer struct {
	Depth     int
	Range     int
	Scanner   int
	Direction int
}

type Layers struct {
	LayersByDepth map[int]*Layer
	Position      int
	Picosecond    int
	Caught        bool
	Severity      int
}

func New(m map[int]int) (ll Layers) {
	ll.LayersByDepth = make(map[int]*Layer, len(m))
	for d, r := range m {
		ll.LayersByDepth[d] = &Layer{d, r, 0, 1}
	}
	ll.Position = -1
	ll.Picosecond = -1
	return
}

func (ll *Layers) AdvancePosition() {
	ll.Position += 1
	if l, ok := ll.LayersByDepth[ll.Position]; ok {
		if l.Scanner == 0 {
			ll.Caught = true
			ll.Severity += l.Depth * l.Range
		}
	}
}

func (ll *Layers) AdvanceScanners() {
	for _, l := range ll.LayersByDepth {
		l.Scanner += l.Direction
		if l.Scanner < 0 {
			l.Scanner, l.Direction = l.Scanner+2, 1
		} else if l.Scanner >= l.Range {
			l.Scanner, l.Direction = l.Scanner-2, -1
		}
	}
	ll.Picosecond += 1
}

func (ll *Layers) MaxDepth() (m int) {
	m = math.MinInt32
	for d := range ll.LayersByDepth {
		if d > m {
			m = d
		}
	}
	return
}

func (ll *Layers) Clone() Layers {
	m := make(map[int]*Layer, len(ll.LayersByDepth))
	for d, l := range ll.LayersByDepth {
		l2 := *l
		m[d] = &l2
	}
	return Layers{m, ll.Position, ll.Picosecond, ll.Caught, ll.Severity}
}
