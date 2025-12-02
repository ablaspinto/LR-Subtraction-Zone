package main

import (
	"LRZ/src/player"
	"fmt"
)

var listOfWinners []map[int]player.Player

func main() {
	for {
		counter, leftOption, rightOption := player.HandleInput()
		var original int = counter
		var left, right = player.CreatePlayers(leftOption, rightOption)
		if player.ProperOption(rightOption) && player.ProperOption(leftOption) {
			fmt.Printf("%d", original)
			for {
				playerMap := make(map[int]player.Player)
				var num, player = player.PlayAgain(left, right, counter)
				playerMap[num] = player
				listOfWinners = append(listOfWinners, playerMap)
				break
			}
			fmt.Printf("%v\n", listOfWinners)
		} else {
			fmt.Printf("Left or Right have Invalid Options\n")
			fmt.Printf("Options must be > 0 and <= 10\n")
		}
	}
}
