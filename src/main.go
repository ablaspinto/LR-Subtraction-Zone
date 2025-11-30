package main

import (
	"LRZ/src/player"
	"fmt"
	"strconv"
)

const Reset string = "\033[0m"
const Red string = "\033[31m"
const Blue string = "\033[34m"

func main() {
	for {
		counter, leftOption, rightOption := player.HandleInput()
		var original int = counter
		left, right := player.CreatePlayers(leftOption, rightOption)
		if player.ProperOption(rightOption) && player.ProperOption(leftOption) {
			fmt.Printf(strconv.Itoa(original) + Blue + " -L-> " + Reset)
			counter = player.Dec(left, counter)
			countString := strconv.Itoa(counter)
			fmt.Printf(countString + Red + "-R->" + Reset)
			counter = player.Dec(right, counter)
			rcountString := strconv.Itoa(counter)
			fmt.Printf(" " + rcountString)

		} else {
			fmt.Printf("Left or Right have Invalid Options\n")
			fmt.Printf("Options must be > 0 and <= 10\n")
		}
	}
}
