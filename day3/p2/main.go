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

	var priorities int
	//three-Elf group
	var lines [3]string

	var j int = 1
	for {
		//read line from file
		_, err := fmt.Fscanf(file, "%s\n", &lines[j-1])
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		if j%3 == 0 {
			j = 0
			for i := 0; i < len(lines[0]); i++ {
				iFirst := strings.Index(lines[1], string(lines[0][i]))
				iSecond := strings.Index(lines[2], string(lines[0][i]))

				if iFirst != -1 && iSecond != -1 {

					//To help prioritize item rearrangement, every item type
					//can be converted to a priority:
					//Lowercase item types a through z have priorities 1 through 26.
					//Uppercase item types A through Z have priorities 27 through 52.
					var priority int

					if lines[0][i] >= 97 && lines[0][i] <= 122 {
						// a equals 97 in decimal at ASCII chart
						// minus -1 because priority index start 1
						priority = int(lines[0][i] - 96)
					} else {
						// A equals 65 in decimal at ASCII chart
						// minus -1 because priority index start 1
						priority = int(lines[0][i] - 64 + 26)
					}
					priorities += priority
					break
				}
			}
		}
		j++
	}
	fmt.Println("The sum of the priorities: ", priorities)
}
