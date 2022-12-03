package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//open file
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var line string
	var priorities int

	for {
		//read line from file
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		var firstHalf string
		var secondHalf string

		firstHalf = line[:len(line)/2]
		secondHalf = line[len(line)/2:]

		for i := 0; i < len(firstHalf); i++ {
			index := strings.Index(secondHalf, string(firstHalf[i]))
			if index != -1 {

				//To help prioritize item rearrangement, every item type can be converted to a priority:
				//Lowercase item types a through z have priorities 1 through 26.
				//Uppercase item types A through Z have priorities 27 through 52.
				var priority int

				if secondHalf[index] >= 97 && secondHalf[index] <= 122 {
					// a equals 97 in decimal at ASCII chart
					// minus -1 because priority index start 1
					priority = int(secondHalf[index] - 96)
				} else {
					// A equals 65 in decimal at ASCII chart
					// minus -1 because priority index start 1
					priority = int(secondHalf[index] - 64 + 26)
				}
				priorities += priority
				break
			}
		}
	}
	fmt.Println("The sum of the priorities: ", priorities)
}
