package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var maxTotal int
	var nth int
	type tElf struct {
		nth      int //index
		calories int //max carrying calories
	}

	var elves = []tElf{}

	for {
		var curTotalCal int = 0
		var curNum int = 0
		// pattern to scan
		_, err := fmt.Fscanf(file, "%d\n", &curNum)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		//fmt.Println("curNum: ", curNum)
		curTotalCal += curNum
		for {
			// pattern to scan
			_, err := fmt.Fscanf(file, "%d\n", &curNum)
			if err != nil {
				if err == io.EOF {
					break
				}
				break
			}
			//fmt.Println("curNum: ", curNum)
			curTotalCal += curNum
		}

		if curTotalCal > maxTotal {
			//fmt.Printf("curTotalCal: %d - maxTotal: %d\n", curTotalCal, maxTotal)
			maxTotal = curTotalCal
		}
		elves = append(elves, tElf{nth, curTotalCal})
		nth++
	}

	//fmt.Println("The Elves : ", elves)
	//fmt.Println("Sort the Elves according to the carrying Calories!")
	// Sort by Calories, keeping original order or equal elements.
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
	//fmt.Println("After sorting process")
	//fmt.Println("The Elves : ", elves)

	//Sum up top three Elves
	var maxTopThreeTotalCal int
	for i := 0; i < 3; i++ {
		maxTopThreeTotalCal += elves[i].calories
	}

	fmt.Println("The top three Elves carrying the most Calories is : ", maxTopThreeTotalCal)
}
