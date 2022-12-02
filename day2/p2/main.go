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

func calculateNextMove(oMove byte, res byte) byte {
	var move byte

	switch res {
	// lose
	case 'X':
		switch oMove {
		// ROCK
		case 'A':
			move = 'Z'
			break
		//PAPER
		case 'B':
			move = 'X'
			break
		//SCISSORS
		case 'C':
			move = 'Y'
			break
		}
		break
	//draw
	case 'Y':
		switch oMove {
		// ROCK
		case 'A':
			move = 'X'
			break
		//PAPER
		case 'B':
			move = 'Y'
			break
		//SCISSORS
		case 'C':
			move = 'Z'
			break
		}
		break
	//win
	case 'Z':
		switch oMove {
		// ROCK
		case 'A':
			move = 'Y'
			break
		//PAPER
		case 'B':
			move = 'Z'
			break
		//SCISSORS
		case 'C':
			move = 'X'
			break
		}
		break
	}

	return move
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
	var res byte

	for {
		_, err := fmt.Fscanf(file, "%c %c\n", &oMove, &res)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		score += calculateScore(oMove, calculateNextMove(oMove, res))
	}

	fmt.Println("Total score: ", score)

}
