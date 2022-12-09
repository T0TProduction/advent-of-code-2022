package main

import (
	"fmt"
	"os"
	"strings"
)

func StringContains(s string, r rune) bool {
	for _, item := range s {
		if item == r {
			return true
		}
	}
	return false
}

func GetPriorityOfRune(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - 'a' + 1
	}
	if r >= 'A' && r <= 'Z' {
		return int(r) - 'A' + 27
	}
	return 0
}

// calculate both solutions
func main() {

	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	allBackpacks := strings.Split(fileContent, "\n")

	sumOfPriorities := 0

	for _, currentBackpack := range allBackpacks {
		firstHalf := currentBackpack[:len(currentBackpack)/2]
		secondHalf := currentBackpack[len(currentBackpack)/2:]

		var commonItem rune

		for _, item := range firstHalf {
			if StringContains(secondHalf, item) {
				commonItem = (item)
				break
			}
		}

		sumOfPriorities += GetPriorityOfRune(commonItem)

	}

	fmt.Println("Sum of priorities:", sumOfPriorities)

}
