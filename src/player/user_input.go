package player

import "fmt"

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
