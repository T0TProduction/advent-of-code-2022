package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	listOfElves := []int{}
	var currentSum int64 = 0

	allNumbers := strings.Split(fileContent, "\n")

	for _, currentNumber := range allNumbers {
		if currentNumber == "" {
			listOfElves = append(listOfElves, int(currentSum))
			currentSum = 0
		}
		amount, _ := strconv.ParseInt(currentNumber, 10, 64)
		currentSum += amount
	}

	sort.Ints(listOfElves)

	highestAmount := listOfElves[len(listOfElves)-1]

	fmt.Println("Highest amount:", highestAmount)

	highestThreeAmounts := listOfElves[len(listOfElves)-3:]

	highestThreeAmountsCombined := 0

	for _, currentAmount := range highestThreeAmounts {
		highestThreeAmountsCombined += currentAmount
	}

	fmt.Println("Combined amount of highest three amount:", highestThreeAmountsCombined)

}
