package p19

import (
	"strings"
)

func Traverse(m []string) (p string, t int) {
	y := 0
	x := strings.Index(m[y], "|")
	dx, dy := 0, 1
	t = 1
	for {
		switch c := m[y][x]; c {
		case '|', '+', '-':
		case ' ':
			return
		default:
			p += string(c)
		}
		if inPath(m, x+dx, y+dy) {
		} else if dy == 0 {
			if inPath(m, x, y+1) {
				dx, dy = 0, 1
			} else if inPath(m, x, y-1) {
				dx, dy = 0, -1
			} else {
				return
			}
		} else if dx == 0 {
			if inPath(m, x+1, y) {
				dx, dy = 1, 0
			} else if inPath(m, x-1, y) {
				dx, dy = -1, 0
			} else {
				return
			}
		} else {
			return
		}
		x, y, t = x+dx, y+dy, t+1
	}
}

func inRange(m []string, x, y int) bool {
	return y < len(m) && x < len(m[y])
}

func inPath(m []string, x, y int) bool {
	return inRange(m, x, y) && m[y][x] != ' '
}
