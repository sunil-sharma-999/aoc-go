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
	cycle int
	count int
	value int
}

func drawInCRT(crt *[6][40]string, rowIndex, colIndex, position int) {
	if colIndex == position || colIndex == position+2 || colIndex == position+1 {
		crt[rowIndex][colIndex] = "#"
	}
}

func main() {
	file, err := os.Open("./2022/day10/input.txt")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	totalCycles := 0
	X := 1
	instructions := []*Instruction{}
	crt := [6][40]string{}

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
			cycle: cycle,
			count: cycle,
			value: value,
		})
	}

	step := 0
	stripePosition := X - 1
	for i := 0; i < totalCycles; i++ {
		instruct := instructions[step]

		if instruct.count == 0 {
			step++
			// register X
			X += instruct.value
			stripePosition = X - 1
			instructions[step].count--
		} else {
			instruct.count--
		}
		if i >= 0 && i < 40 {
			drawInCRT(&crt, 0, i, stripePosition)
		} else if i >= 40 && i < 80 {
			drawInCRT(&crt, 1, i-40, stripePosition)
		} else if i >= 80 && i < 120 {
			drawInCRT(&crt, 2, i-80, stripePosition)
		} else if i >= 120 && i < 160 {
			drawInCRT(&crt, 3, i-120, stripePosition)
		} else if i >= 160 && i < 200 {
			drawInCRT(&crt, 4, i-160, stripePosition)
		} else if i >= 200 && i < 240 {
			drawInCRT(&crt, 5, i-200, stripePosition)
		}

	}
	for _, arr := range crt {
		for _, j := range arr {
			if j == "#" {
				fmt.Print(j)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
