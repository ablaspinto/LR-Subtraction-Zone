package main

import (
	"LRZ/src/player"
	"fmt"
)

var leftGoingFirst []map[int]player.Player
var rightGoingFirst []map[int]player.Player

func main() {
	counters, allPlayerMoves := player.GetAllNumbers()
	leftPlayerMoves := allPlayerMoves["L"]
	rightPlayerMoves := allPlayerMoves["R"]
	// crosss reference maps
	for i := 0; i < len(counters); i++ {
		currLeftMove := leftPlayerMoves[i]
		currRightMove := rightPlayerMoves[i]
		countNumber, lgfWinners := player.PlayAgain(currLeftMove, currRightMove, counters[i])
		countNumber, rgfWinners := player.PlayAgain(currRightMove, currLeftMove, counters[i])
		m1 := make(map[int]player.Player)
		m2 := make(map[int]player.Player)
		m1[countNumber] = lgfWinners
		m2[countNumber] = rgfWinners
		leftGoingFirst = append(leftGoingFirst, m1)
		rightGoingFirst = append(rightGoingFirst, m2)
	}
	fmt.Printf("%v\n", rightGoingFirst)
	fmt.Printf("%v\n", leftGoingFirst)

}
