package main

import (
	"fmt"
	"io"
	"os"
)

var gamePoints = map[string]int{
	"LOSE": 0,
	"DRAW": 3,
	"WIN":  6,
}

var movePoints = map[string]int{
	"ROCK":     1,
	"PAPER":    2,
	"SCISSORS": 3,
}

func calculateScore(oMove byte, myMove byte) int {
	var score int

	switch oMove {
	// rock
	case 'A':
		switch myMove {
		case 'X':
			score = gamePoints["DRAW"] + movePoints["ROCK"]
			break
		case 'Y':
			score = gamePoints["WIN"] + movePoints["PAPER"]
			break
		case 'Z':
			score = gamePoints["LOSE"] + movePoints["SCISSORS"]
			break
		}
		break
	// paper
	case 'B':
		switch myMove {
		case 'X':
			score = gamePoints["LOSE"] + movePoints["ROCK"]
			break
		case 'Y':
			score = gamePoints["DRAW"] + movePoints["PAPER"]
			break
		case 'Z':
			score = gamePoints["WIN"] + movePoints["SCISSORS"]
			break
		}
		break
	// scissors
	case 'C':
		switch myMove {
		case 'X':
			score = gamePoints["WIN"] + movePoints["ROCK"]
			break
		case 'Y':
			score = gamePoints["LOSE"] + movePoints["PAPER"]
			break
		case 'Z':
			score = gamePoints["DRAW"] + movePoints["SCISSORS"]
			break
		}
		break
	}

	return score
}

func main() {
	// open file
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//total score
	var score int
	var oMove byte
	var myMove byte

	for {
		_, err := fmt.Fscanf(file, "%c %c\n", &oMove, &myMove)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		score += calculateScore(oMove, myMove)
	}

	fmt.Println("Total score: ", score)

}
