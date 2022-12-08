package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = "A"
	Paper    = "B"
	Scissors = "C"
)

const (
	MyRock     = "X"
	MyPaper    = "Y"
	MyScissors = "Z"
)

const (
	RockScore     = 1
	PaperScore    = 2
	ScissorsScore = 3
)

const (
	WinScore  = 6
	DrawScore = 3
	LoseScore = 0
)

func MapMyMoveToNormalizedMove(myMove string) string {
	switch myMove {
	case MyRock:
		return Rock
	case MyPaper:
		return Paper
	case MyScissors:
		return Scissors
	}
	return ""
}

func MapMyMoveToMoveScore(myMove string) int {
	switch myMove {
	case Rock:
		return RockScore
	case Paper:
		return PaperScore
	case Scissors:
		return ScissorsScore
	}
	return 0
}

func MatchOutcome(myMove string, enemyMove string) int {
	if myMove == enemyMove {
		return DrawScore
	}
	if myMove == Rock {
		if enemyMove == Paper {
			return LoseScore
		}
		if enemyMove == Scissors {
			return WinScore
		}
	}
	if myMove == Paper {
		if enemyMove == Rock {
			return WinScore
		}
		if enemyMove == Scissors {
			return LoseScore
		}
	}
	if myMove == Scissors {
		if enemyMove == Rock {
			return LoseScore
		}
		if enemyMove == Paper {
			return WinScore

		}
	}

	return 0
}

// Part two helpers

const (
	NeedToLose = "X"
	NeedToDraw = "Y"
	NeedToWin  = "Z"
)

func GetMyMoveFromOutcome(enemyMove string, outcome string) string {
	if outcome == NeedToDraw {
		return enemyMove
	}

	switch enemyMove {
	case Rock:
		if outcome == NeedToLose {
			return Scissors
		}
		if outcome == NeedToWin {
			return Paper
		}
	case Paper:
		if outcome == NeedToLose {
			return Rock
		}

		if outcome == NeedToWin {
			return Scissors
		}
	case Scissors:
		if outcome == NeedToLose {
			return Paper
		}

		if outcome == NeedToWin {
			return Rock
		}
	}
	return ""
}

// calculate both solutions
func main() {

	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	allMatches := strings.Split(fileContent, "\n")

	sumOfPoints := 0
	sumOfPointsBasedOnOutcome := 0

	for _, currentMatch := range allMatches {
		enemyMove := currentMatch[0:1]
		myMove := MapMyMoveToNormalizedMove(currentMatch[2:3])

		outcome := MatchOutcome(myMove, enemyMove) + MapMyMoveToMoveScore(myMove)
		sumOfPoints += outcome

		// basedd on outcome
		outcomeWanted := currentMatch[2:3]
		myMoveBasedOnOutcome := GetMyMoveFromOutcome(enemyMove, outcomeWanted)
		outcomeBasedOnOutcome := MatchOutcome(myMoveBasedOnOutcome, enemyMove) + MapMyMoveToMoveScore(myMoveBasedOnOutcome)
		sumOfPointsBasedOnOutcome += outcomeBasedOnOutcome
	}

	fmt.Println("Outcome part one:", sumOfPoints)
	fmt.Println("Outcome part two:", sumOfPointsBasedOnOutcome)

}
