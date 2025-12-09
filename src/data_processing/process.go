package dataprocessing

import (
	"LRZ/src/player"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type GameResult struct {
	Counter         int
	LeftOption      int
	RightOption     int
	LeftGoingFirst  player.Player
	RightGoingFirst player.Player
}

func ClassifyPositions(results []GameResult) {
	file, err := os.Create("slrz_results.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"L Option", "R Option", "Counter", "Outcome Class", "LGF Winner", "RGF Winner"})

	outcomeCounts := map[string]int{"N": 0, "P": 0, "L": 0, "R": 0}

	for _, result := range results {
		lgf := result.LeftGoingFirst.Name
		rgf := result.RightGoingFirst.Name

		var outcome string
		if lgf == "L" && rgf == "R" {
			outcome = "N"
		} else if lgf == "L" && rgf == "L" {
			outcome = "L"
		} else if lgf == "R" && rgf == "R" {
			outcome = "R"
		} else if lgf == "R" && rgf == "L" {
			outcome = "P"
		}

		outcomeCounts[outcome]++

		writer.Write([]string{
			strconv.Itoa(result.LeftOption),
			strconv.Itoa(result.RightOption),
			strconv.Itoa(result.Counter),
			outcome,
			lgf,
			rgf,
		})
	}

	fmt.Println("Results saved to slrz_results.csv")
	fmt.Println("\n--- Summary ---")
	fmt.Printf("Total games: %d\n", len(results))
	fmt.Printf("N (Next wins):     %d\n", outcomeCounts["N"])
	fmt.Printf("P (Previous wins): %d\n", outcomeCounts["P"])
	fmt.Printf("L (Left wins):     %d\n", outcomeCounts["L"])
	fmt.Printf("R (Right wins):    %d\n", outcomeCounts["R"])
}
