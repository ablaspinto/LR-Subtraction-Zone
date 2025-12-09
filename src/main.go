package main

import (
	"LRZ/src/data_processing"
	"LRZ/src/player"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	counters, allPlayerMoves := player.GetAllNumbers()
	leftPlayerMoves := allPlayerMoves["L"]
	rightPlayerMoves := allPlayerMoves["R"]

	var allResults []dataprocessing.GameResult
	for i := 0; i < len(leftPlayerMoves); i++ {
		currLeft := leftPlayerMoves[i]
		currRight := rightPlayerMoves[i]
		for j := 0; j < len(counters); j++ {
			counter := counters[j]
			_, lgfWinner := player.PlayAgain(currLeft, currRight, counter)
			_, rgfWinner := player.PlayAgain(currRight, currLeft, counter)
			allResults = append(allResults, dataprocessing.GameResult{
				Counter:         counter,
				LeftOption:      currLeft.Option,
				RightOption:     currRight.Option,
				LeftGoingFirst:  lgfWinner,
				RightGoingFirst: rgfWinner,
			})
		}
	}
	dataprocessing.ClassifyPositions(allResults)
	duration := time.Since(start)
	fmt.Printf("TOTAL DURATION OF PROGRAM %v\n", duration.Seconds())
}
