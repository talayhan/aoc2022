package main

import (
	"fmt"
	"io"
	"os"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//elves pair
	var x1, x2, x3, x4 int
	//the number of pairs matched the rule
	var count int
	for {
		_, err := fmt.Fscanf(file, "%d-%d,%d-%d\n", &x1, &x2, &x3, &x4)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		firstPair := mapset.NewSet[int]()
		for i := x1; i <= x2; i++ {
			firstPair.Add(i)
		}

		secondPair := mapset.NewSet[int]()
		for i := x3; i <= x4; i++ {
			secondPair.Add(i)
		}

		intersect := secondPair.Intersect(firstPair)
		if intersect.Cardinality() != 0 {
			count++
		}
	}
	fmt.Println("The number of pairs do the ranges overlap is ", count)
}
