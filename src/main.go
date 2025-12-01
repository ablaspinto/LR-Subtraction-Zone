package main

import (
	"LRZ/src/player"
	"fmt"
)

const Reset string = "\033[0m"

func main() {
	for {
		counter, leftOption, rightOption := player.HandleInput()
		var left, right player.Player
		var original int = counter
		left, right = player.CreatePlayers(leftOption, rightOption)
		if player.ProperOption(rightOption) && player.ProperOption(leftOption) {
			fmt.Printf("%d", original)
			for {
				counter = player.Dec(left, counter)
				if counter < 0 {
					fmt.Println("Right Wins!\n")
					break
				}
				fmt.Printf(player.PrintFunction(counter, left))
				counter = player.Dec(right, counter)
				if counter < 0 {
					fmt.Println("\nLeft Wins!\n")
					break
				}
				fmt.Printf(player.PrintFunction(counter, right))
			}
		} else {
			fmt.Printf("Left or Right have Invalid Options\n")
			fmt.Printf("Options must be > 0 and <= 10\n")
		}
	}
}
