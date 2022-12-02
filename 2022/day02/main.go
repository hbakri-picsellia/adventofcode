package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

const (
	Rock             string = "A"
	Paper                   = "B"
	Scissors                = "C"
	RockResponse            = "X"
	PaperResponse           = "Y"
	ScissorsResponse        = "Z"
	Lose                    = "X"
	Draw                    = "Y"
	Win                     = "Z"
)

func getShapeScore(response string) (score int) {
	switch response {
	case RockResponse:
		score = 1
	case PaperResponse:
		score = 2
	case ScissorsResponse:
		score = 3
	}
	return score
}

func RockPaperScissors(player1 string, player2 string) (score int) {
	switch player1 {
	case Rock:
		switch player2 {
		case RockResponse:
			score += 3
		case PaperResponse:
			score += 6
		case ScissorsResponse:
			score += 0
		}
	case Paper:
		switch player2 {
		case RockResponse:
			score += 0
		case PaperResponse:
			score += 3
		case ScissorsResponse:
			score += 6
		}
	case Scissors:
		switch player2 {
		case RockResponse:
			score += 6
		case PaperResponse:
			score += 0
		case ScissorsResponse:
			score += 3
		}
	}
	score += getShapeScore(player2)
	return score
}

func findResponse(player1 string, score string) (response string) {
	switch player1 {
	case Rock:
		switch score {
		case Lose:
			response = ScissorsResponse
		case Draw:
			response = RockResponse
		case Win:
			response = PaperResponse
		}
	case Paper:
		switch score {
		case Lose:
			response = RockResponse
		case Draw:
			response = PaperResponse
		case Win:
			response = ScissorsResponse
		}
	case Scissors:
		switch score {
		case Lose:
			response = PaperResponse
		case Draw:
			response = ScissorsResponse
		case Win:
			response = RockResponse
		}
	}
	return response
}

func step1(input string) (points int) {
	rounds := strings.Split(input, "\n")
	for _, value := range rounds {
		round := strings.Split(value, " ")
		points += RockPaperScissors(round[0], round[1])
	}
	return points
}

func step2(input string) (points int) {
	rounds := strings.Split(input, "\n")
	for _, value := range rounds {
		round := strings.Split(value, " ")
		response := findResponse(round[0], round[1])
		points += RockPaperScissors(round[0], response)
	}
	return points
}

func main() {
	fmt.Println("--- Day 2: Rock Paper Scissors ---")

	example := utils.ParseFileToString("2022/day02/example.txt")
	utils.AssertEqual(step1(example), 15, "example step1")
	utils.AssertEqual(step2(example), 12, "example step2")

	input := utils.ParseFileToString("2022/day02/input.txt")
	utils.AssertEqual(step1(input), 15632, "step1")
	utils.AssertEqual(step2(input), 14416, "step2")
}
