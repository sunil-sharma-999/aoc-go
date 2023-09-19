package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var CYCLES = map[string]int{
	"addx": 2,
	"noop": 1,
}

type Instruction struct {
	name  string
	cycle int
	count int
	value int
}

func main() {

	totalCycles := 0
	strength := 0

	X := 1

	file, err := os.Open("./2022/day10/input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	instructions := []*Instruction{}

	for scanner.Scan() {
		text := scanner.Text()
		instruction := strings.Fields(text)
		cycle := CYCLES[instruction[0]]

		totalCycles += cycle

		var value int

		if len(instruction) == 2 {
			value, err = strconv.Atoi(instruction[1])
			if err != nil {
				log.Fatal(err)
			}
		}

		instructions = append(instructions, &Instruction{
			name:  instruction[0],
			cycle: cycle,
			count: cycle,
			value: value,
		})

	}

	step := 0
	for i := 0; i < totalCycles; i++ {
		instruct := instructions[step]

		if instruct.count == 0 {
			step++
			X += instruct.value
			instructions[step].count--
		} else {
			instruct.count--
		}

		// 20th, 60th, 100th, 140th, 180th, and 220th cycles
		switch i + 1 {
		case 20, 60, 100, 140, 180, 220:
			strength += ((i + 1) * X)
		}

	}
	fmt.Println(strength)

}
