package p

import (
	"bytes"
	"container/ring"
)

func P1(q int) (s string) {
	n := 2
	first := ring.New(n)
	first.Value = 3
	second := first.Next()
	second.Value = 7
	last := second
	add := func(v int) {
		n++
		l := ring.New(1)
		l.Value = v
		last.Link(l)
		last = l
		if n > q && len(s) < 10 {
			s += string('0' + v)
		}
	}
	for len(s) != 10 {
		sum := first.Value.(int) + second.Value.(int)
		d1, d2 := sum/10, sum%10
		if d1 > 0 {
			add(d1)
		}
		add(d2)
		first = first.Move(first.Value.(int) + 1)
		second = second.Move(second.Value.(int) + 1)
	}
	return
}

func P2(q int, m string) int {
	board := []byte{}
	board = append(board, 3, 7)
	i, j, mm := 0, 1, []byte{}
	for i := range m {
		mm = append(mm, m[i]-'0')
	}
	for {
		sum := board[i] + board[j]
		d1, d2 := sum/10, sum%10
		if d1 > 0 {
			board = append(board, byte(d1))
			if len(board) >= len(m) {
				if bytes.Compare(board[len(board)-len(m):], mm) == 0 {
					return len(board) - len(m)
				}
			}
		}
		board = append(board, byte(d2))
		if len(board) >= len(m) {
			if bytes.Compare(board[len(board)-len(m):], mm) == 0 {
				return len(board) - len(m)
			}
		}
		i = (i + int(board[i]) + 1) % len(board)
		j = (j + int(board[j]) + 1) % len(board)
	}
}
