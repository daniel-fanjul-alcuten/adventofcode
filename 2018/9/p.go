package p

import (
	"container/ring"
)

func P1(numPlayers int, lastMarble int) (highestScore int) {
	currentPlayer := ring.New(numPlayers)
	currentMarble := ring.New(1)
	currentMarble.Value = 0
	for m := 1; m <= lastMarble; m++ {
		if m%23 == 0 {
			currentMarble = currentMarble.Move(-8)
			score := 0
			if currentPlayer.Value != nil {
				score = currentPlayer.Value.(int)
			}
			score += m + currentMarble.Unlink(1).Value.(int)
			currentPlayer.Value = score
			if score > highestScore {
				highestScore = score
			}
			currentMarble = currentMarble.Move(1)
		} else {
			currentMarble = currentMarble.Move(1)
			currentMarble = currentMarble.Link(ring.New(1))
			currentMarble = currentMarble.Move(-1)
			currentMarble.Value = m
		}
		currentPlayer = currentPlayer.Move(1)
	}
	return
}
