package p

import (
	"math"
)

func PowerLevel(x, y, gridSerialNumber int) (powerLevel int) {
	rackID := 10 + x
	powerLevel = rackID * y
	powerLevel += gridSerialNumber
	powerLevel *= rackID
	powerLevel = (powerLevel / 100) % 10
	powerLevel -= 5
	return
}

func P1(lastX, lastY, gridSerialNumber int) (bestX, bestY, powerLevel int) {
	powerLevel = math.MinInt32
	for x := 1; x < lastX-1; x++ {
		for y := 1; y < lastY-1; y++ {
			pl := 0
			for ix := 0; ix < 3; ix++ {
				for iy := 0; iy < 3; iy++ {
					pl += PowerLevel(x+ix, y+iy, gridSerialNumber)
				}
			}
			if pl > powerLevel {
				bestX, bestY, powerLevel = x, y, pl
			}
		}
	}
	return
}

func P2(lastXY, gridSerialNumber int) (bestX, bestY, bestZ, powerLevel int) {
	pl := map[int]map[int]int{}
	for x := 1; x < lastXY-1; x++ {
		plx := map[int]int{}
		pl[x] = plx
		for y := 1; y < lastXY-1; y++ {
			plx[y] = PowerLevel(x, y, gridSerialNumber)
		}
	}
	powerLevel = math.MinInt32
	zpl := map[int]map[int]map[int]int{}
	stats1 := 0
	stats2 := 0
	for z := 1; z <= lastXY; z++ {
		xpl := map[int]map[int]int{}
		zpl[z] = xpl
		for x := lastXY - z + 1; x >= 1; x-- {
			ypl := map[int]int{}
			xpl[x] = ypl
			for y := lastXY - z + 1; y >= 1; y-- {
				s := 0
				if s2, ok := zpl[z-1][x+1][y+1]; !ok {
					for ix := 0; ix < z; ix++ {
						for iy := 0; iy < z; iy++ {
							s += pl[x+ix][y+iy]
						}
					}
					stats1++
				} else {
					s += s2
					for ix := 0; ix < z; ix++ {
						s += pl[x+ix][y]
					}
					for iy := 1; iy < z; iy++ {
						s += pl[x][y+iy]
					}
					stats2++
				}
				ypl[y] = s
				if s > powerLevel {
					bestX, bestY, bestZ, powerLevel = x, y, z, s
				}
			}
		}
	}
	return
}
