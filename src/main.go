package main

import (
	"LRZ/src/player"
	"fmt"
)

var listOfWinners []map[int]player.Player

func main() {
	counters, allPlayerMoves := player.GetAllNumbers()
	fmt.Printf("%v", counters)
	fmt.Printf("%v", allPlayerMoves)
}
