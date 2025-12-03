package player

import (
	"fmt"
	"strconv"
)

const Reset string = "\033[0m"

func PrintFunction(number int, p Player) string {
	var stringStatement string
	switch p.name {
	case "L":
		stringStatement = Reset + p.color + " -L-> " + Reset + strconv.Itoa(number)
	case "R":
		stringStatement = Reset + p.color + " -R-> " + Reset + strconv.Itoa(number)
	}
	return stringStatement
}

type Player struct {
	name   string
	option int
	color  string
}

func CreatePlayers(leftOption int, rightOption int) (Player, Player) {
	var left = Player{
		name:   "L",
		option: leftOption,
		color:  "\033[34m",
	}
	var right = Player{
		name:   "R",
		option: rightOption,
		color:  "\033[31m",
	}
	return left, right

}

func detZone(num int) int {
	return num % 5
}

func leftZone(p Player, modulusNumber int) bool {
	return p.name == "L" && (modulusNumber == 3 || modulusNumber == 4)
}

func rightZone(p Player, modulusNumber int) bool {
	return p.name == "R" && (modulusNumber == 0 || modulusNumber == 1)
}

func Dec(player Player, stack int) int {
	var temp int = stack
	temp = temp - player.option
	if temp == 0 {
		return temp
	}
	var possibleSub = player.option * 2
	var modulusNumber = detZone(temp)
	if leftZone(player, modulusNumber) {
		stack = stack - possibleSub
	} else if rightZone(player, modulusNumber) {
		stack = stack - possibleSub
	} else {
		stack = stack - player.option
	}
	return stack
}

func PlayAgain(l Player, r Player, counter int) (int, Player) {
	var ogCounter = counter
	for {
		counter = Dec(l, counter)
		if counter < 0 {
			fmt.Println("\nRight Wins!")
			return ogCounter, r
		}
		var lword string = PrintFunction(counter, l)
		fmt.Printf("%s", lword)
		counter = Dec(r, counter)
		if counter < 0 {
			fmt.Println("\nLeft Wins!")
			return ogCounter, l
		}
		var rword string = PrintFunction(counter, r)
		fmt.Printf("%s", rword)
	}
}
