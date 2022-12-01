package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var maxTotal int
	//var nth int
	//keep list double array

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
	}

	fmt.Println("The Elf carrying the most Calories is : ", maxTotal)
}
