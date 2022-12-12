package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Grid struct {
	grid   [][]int
	width  int
	length int
}

var maxScenicScore int

func (g Grid) isVisible(row, column int) int {
	var leftCount, rightCount, upCount, bottomCount int

	// left
	for i := column - 1; i >= 0; i-- {

		if g.grid[row][column] > g.grid[row][i] {
			leftCount++
		}

		if g.grid[row][column] <= g.grid[row][i] {
			leftCount++
			break
		}
	}

	// right
	for i := column + 1; i < g.length; i++ {
		if g.grid[row][column] > g.grid[row][i] {
			rightCount++
		}

		if g.grid[row][column] <= g.grid[row][i] {
			rightCount++
			break
		}
	}

	// up
	for i := row - 1; i >= 0; i-- {
		if g.grid[row][column] > g.grid[i][column] {
			upCount++
		}
		if g.grid[row][column] <= g.grid[i][column] {
			upCount++
			break
		}
	}

	// down
	for i := row + 1; i < g.width; i++ {
		if g.grid[row][column] > g.grid[i][column] {
			bottomCount++
		}
		if g.grid[row][column] <= g.grid[i][column] {
			bottomCount++
			break
		}
	}

	isVisibleCount := leftCount * rightCount * upCount * bottomCount
	if maxScenicScore < isVisibleCount {
		maxScenicScore = isVisibleCount
	}

	return maxScenicScore
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid Grid
	grid.grid = make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var length int

		tmpArr := make([]int, len(line))
		for length = 0; length < len(line); length++ {
			tmpArr[length] = int(line[length]) - '0'
		}
		grid.grid = append(grid.grid, tmpArr)
		//TODO need one time assignment
		grid.length = length
		grid.width = length
	}

	for i := 1; i < grid.width; i++ {
		for j := 1; j < grid.length; j++ {
			// skip edges
			if i == grid.width-1 {
				continue
			}
			if j == grid.length-1 {
				continue
			}

			grid.isVisible(i, j)
		}
	}

	fmt.Println("Max scenic score: ", maxScenicScore)
}
