package player

import (
	"fmt"
)

func ProperOption(num int) bool {
	return num > 0 && num <= 10
}

func HandleInput() (int, int, int) {
	var counters int
	var leftOption int
	var rightOption int
	fmt.Printf("How large would you like the counters to be?: \n")
	fmt.Scan(&counters)
	fmt.Printf("What would you like Left's option to be?\n")
	fmt.Scan(&leftOption)
	fmt.Printf("What would you like Right's option to be?\n")
	fmt.Scan(&rightOption)
	return counters, leftOption, rightOption
}

func GetAllNumbers() ([100]int, map[string][]Player) {
	var intArr [100]int
	var playerMoves = make(map[string][]Player)
	playerMoves["L"] = []Player{}
	playerMoves["R"] = []Player{}
	var j int = 0
	for i := 1; i < 101; i++ {
		intArr[j] = i
		j++
	}
	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			left, right := CreatePlayers(i, j)
			playerMoves["L"] = append(playerMoves["L"], left)
			playerMoves["R"] = append(playerMoves["R"], right)
		}
	}
	return intArr, playerMoves
}
