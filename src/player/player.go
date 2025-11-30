package player

type Player struct {
	name       string
	option     int
	playerWins []int
	zone       [2]int
}

var lz = [2]int{0, 1}
var rz = [2]int{3, 4}

func CreatePlayers(leftOption int, rightOption int) (Player, Player) {
	var left = Player{
		name:       "L",
		option:     leftOption,
		playerWins: make([]int, 0),
		zone:       lz,
	}
	var right = Player{
		name:       "R",
		option:     rightOption,
		playerWins: make([]int, 0),
		zone:       rz,
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
