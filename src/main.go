package src

import (
	"LRZ/src/player"
	"fmt"
)

func main() {
	for {
		counter, leftOption, rightOption := player.HandleInput()
		left, right := player.CreatePlayers(leftOption, rightOption)
		if player.ProperOption(rightOption) && player.ProperOption(leftOption) {
			player.Dec(left, counter)
			// logic
		} else {
			fmt.Printf("Left or Right have Invalid Options\n")
			fmt.Printf("Options must be > 0 and <= 10\n")
		}
	}
}
