package player

type Player struct {
	option     int
	playerWins []int
	zone       [2]int
}

var lz = [2]int{0, 1}
var rz = [2]int{3, 4}

func CreatePlayers(leftOption int, rightOption int) (Player, Player) {
	var left = Player{
		option:     leftOption,
		playerWins: make([]int, 0),
		zone:       lz,
	}
	var right = Player{
		option:     rightOption,
		playerWins: make([]int, 0),
		zone:       rz,
	}
	return left, right

}

func detZone(num int) int {
	return num % 5
}

func rightZone(number int) bool {
	return number == 3 || number == 4
}
func leftZone(number int) bool {
	return number == 1 || number == 0
}

func determineLR(player Player) (int, int) {
	return player.zone[0], player.zone[1]
}

func Dec(player Player, stack int) int {
	var temp int = stack
	temp = temp - player.option
	var possibleSub = player.option * 2
	var modulusNumber = detZone(temp)
}
