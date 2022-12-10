package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type Instruction struct {
	amount, from, to int
}

type Stack struct {
	items []byte
}

func (s *Stack) PopRef() (elem byte) {
	ret := s.items[len(s.items)-1]

	if len(s.items) > 0 {
		s.items = s.items[:len(s.items)-1]
	}

	return ret
}

func (s *Stack) AddRef(elem byte) {
	s.items = append(s.items, elem)
}

func (s Stack) TryPop() (elem byte) {
	var ret byte

	if len(s.items) > 0 {
		ret = s.items[len(s.items)-1]
	}

	return ret
}

func (s Stack) Pop() (elem byte, ss Stack) {
	ret := s.items[len(s.items)-1]

	if len(s.items) > 0 {
		s.items = s.items[:len(s.items)-1]
	}

	return ret, s
}

func (s Stack) Add(elem byte) (ss Stack) {
	s.items = append(s.items, elem)
	return s
}

func (s Stack) TrimSuffix(elem byte) (ss Stack) {
	s.items = bytes.TrimSuffix(s.items, []byte{elem})

	return s
}

func main() {
	stacks := map[int]Stack{}
	instructions := []Instruction{}

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// read&fill stacks
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		for i := 1; i < len(line); i += 4 {
			var stack Stack
			if line[i] >= 'A' && line[i] <= 'Z' {
				stackID := (i / 4) + 1
				//fmt.Printf("Item: %c stackID: %d\n", line[i], stackID)
				stack.AddRef(line[i])
				// _ dummy assignment
				if _, ok := stacks[stackID]; ok {
					s := stacks[stackID]
					s.AddRef(line[i])
					stacks[stackID] = s
				} else {
					stacks[stackID] = stack
				}
			}
		}
	}

	//reverse stacks
	for k := 1; k < len(stacks)+1; k++ {
		for i, j := 0, len(stacks[k].items)-1; i < j; i, j = i+1, j-1 {
			stacks[k].items[i], stacks[k].items[j] = stacks[k].items[j], stacks[k].items[i]
		}
	}

	// read&fill instructions
	for scanner.Scan() {
		var ins Instruction
		line := scanner.Text()
		fmt.Sscanf(line, "move %d from %d to %d", &ins.amount, &ins.from, &ins.to)
		instructions = append(instructions, ins)
	}

	//run instructions
	for i := 0; i < len(instructions); i++ {
		amount := instructions[i].amount
		from := instructions[i].from
		to := instructions[i].to

		var tempStack Stack
		for j := 0; j < amount; j++ {
			var _elem byte
			//pop element
			_elem, stacks[from] = stacks[from].Pop()
			//add element
			tempStack = tempStack.Add(_elem)
			//stacks[to] = stacks[to].Add(_elem)
		}
		for j := 0; j < amount; j++ {
			var _elem byte
			//pop element
			_elem, tempStack = tempStack.Pop()
			//add element
			stacks[to] = stacks[to].Add(_elem)
		}
	}

	fmt.Printf("After the rearrangement procedure completes, what crate ends up on top of each stack?\n")
	//crate will end up on top of each stack
	for i := 1; i < len(stacks)+1; i++ {
		fmt.Printf("%c", stacks[i].TryPop())
	}
	fmt.Println("")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
