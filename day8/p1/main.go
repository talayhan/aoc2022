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

func (g Grid) isVisible(row, column int) bool {

	isVisible := true
	// left
	for i := column - 1; i >= 0; i-- {
		if g.grid[row][i] >= g.grid[row][column] {
			//not visible
			isVisible = false
			break
		}
	}
	if isVisible {
		return true
	}

	isVisible = true
	// right
	for i := column + 1; i < g.length; i++ {
		if g.grid[row][i] >= g.grid[row][column] {
			//not visible
			isVisible = false
			break
		}
	}
	if isVisible {
		return true
	}

	isVisible = true
	// up
	for i := row - 1; i >= 0; i-- {
		if g.grid[i][column] >= g.grid[row][column] {
			//not visible
			isVisible = false
			break
		}
	}
	if isVisible {
		return true
	}

	isVisible = true
	// down
	for i := row + 1; i < g.width; i++ {
		if g.grid[i][column] >= g.grid[row][column] {
			//not visible
			isVisible = false
			break
		}
	}
	if isVisible {
		return true
	}

	return isVisible
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

	visibleCount := (2*grid.length + 2*grid.width) - 4

	internalVisibleCount := 0
	for i := 1; i < grid.width; i++ {
		for j := 1; j < grid.length; j++ {
			// skip edges
			if i == grid.width-1 {
				continue
			}
			if j == grid.length-1 {
				continue
			}

			ret := grid.isVisible(i, j)
			if ret {
				internalVisibleCount++
			}

		}
	}

	visibleCount += internalVisibleCount
	fmt.Println("Visible count: ", visibleCount)
}
