package p

import (
	"log"
	"strconv"
	"strings"
)

type Position struct {
	X, Y int
}

func P1(width, height int, inputs []string) (output int) {
	screen := map[Position]bool{}
	for in, input := range inputs {
		log.Println(in, input)
		if strings.HasPrefix(input, "rect ") {
			i := strings.Index(input, " ")
			j := strings.Index(input, "x")
			ix, err := strconv.Atoi(input[i+1 : j])
			if err != nil {
				panic(err)
			}
			iy, err := strconv.Atoi(input[j+1:])
			if err != nil {
				panic(err)
			}
			for x := 0; x < ix; x++ {
				for y := 0; y < iy; y++ {
					p := Position{x, y}
					screen[p] = true
				}
			}
		} else if strings.HasPrefix(input, "rotate row ") {
			split := strings.Split(input, " ")
			iy, err := strconv.Atoi(split[2][2:])
			if err != nil {
				panic(err)
			}
			ix, err := strconv.Atoi(split[4])
			if err != nil {
				panic(err)
			}
			row := make([]bool, width)
			for x := 0; x < width; x++ {
				row[(x+ix)%width] = screen[Position{x, iy}]
			}
			for x := 0; x < width; x++ {
				screen[Position{x, iy}] = row[x]
			}
		} else if strings.HasPrefix(input, "rotate column ") {
			split := strings.Split(input, " ")
			ix, err := strconv.Atoi(split[2][2:])
			if err != nil {
				panic(err)
			}
			iy, err := strconv.Atoi(split[4])
			if err != nil {
				panic(err)
			}
			col := make([]bool, height)
			for y := 0; y < height; y++ {
				col[(y+iy)%height] = screen[Position{ix, y}]
			}
			for y := 0; y < height; y++ {
				screen[Position{ix, y}] = col[y]
			}
		} else {
			panic(input)
		}
		for y := 0; y < height; y++ {
			s := ""
			for x := 0; x < width; x++ {
				if screen[Position{x, y}] {
					s += "#"
				} else {
					s += "."
				}
			}
			log.Println(s)
		}
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if screen[Position{x, y}] {
				output++
			}
		}
	}
	return
}
