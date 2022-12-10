package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isUnique(str string) (res bool) {
	res = true

	for i := 0; i < len(str)-1; i++ {
		rest := str[i+1:]

		index := strings.Index(rest, string(str[i]))
		if index != -1 {
			res = false
			break
		}
	}

	return res
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			fourChar := line[i : i+14]
			if isUnique(fourChar) {
				fmt.Println("Marker start: ", fourChar, " i+14: ", i+14)
				break
			}
		}
	}
}
