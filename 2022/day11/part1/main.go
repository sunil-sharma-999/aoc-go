package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func doOperation(operand string, left, right int) int {
	switch operand {
	case "+":
		return left + right
	case "*":
		return left * right
	case "-":
		return left - right
	case "/":
		return left / right
	default:
		log.Fatal("Invalid operand")
		return 0
	}
}

type Monkey struct {
	items        []int
	divisibleBy  int
	ifTrue       int
	ifFalse      int
	operation    string
	text         string
	inspectCount int
}

func (m *Monkey) init(monkeys *[]Monkey) {
	textArr := strings.Split(m.text, "\n")
	for i := 0; i < 6; i++ {
		switch i {
		case 1:
			// set items
			items := strings.Split(textArr[i], ":")
			line := strings.TrimSpace(items[1])
			if line == "" {
				continue
			}
			for _, j := range strings.Split(line, ",") {
				num, err := strconv.Atoi(strings.TrimSpace(j))
				if err != nil {
					log.Fatal(err)
				}
				m.items = append(m.items, num)
			}
		case 2:
			items := strings.Split(textArr[i], ":")
			m.operation = strings.TrimSpace(items[1])
		case 3, 4, 5:
			items := strings.Fields(textArr[i])
			num, err := strconv.Atoi(items[len(items)-1])
			if err != nil {
				log.Fatal(err)
			}
			if i == 3 {
				m.divisibleBy = num
			}
			if i == 4 {
				m.ifTrue = num
			}
			if i == 5 {
				m.ifFalse = num
			}
		}
	}

}

func getWorryLevel(num int) int {
	return num / 3
}

func main() {

	path := flag.String("input", "./2022/day11/input.txt", "Input file path")
	rounds := flag.Int("rounds", 20, "Number of rounds")
	flag.Parse()

	file, err := os.Open(*path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	monkeys := []Monkey{}

	lineCounter := 0
	instructionText := ""

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if lineCounter == 5 {
			// create monkey struct and push to array
			instructionText += strings.TrimSpace(text)
			monkeys = append(monkeys, Monkey{text: instructionText})

			// reset
			instructionText = ""
			lineCounter = 0
		} else {
			instructionText += strings.TrimSpace(text) + "\n"
			lineCounter += 1
		}
	}

	for i := 0; i < len(monkeys); i++ {
		monkeys[i].init(&monkeys)
	}

	for i := 0; i < *rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			m := &monkeys[j]

			for _, item := range m.items {
				operation := strings.Fields(m.operation)
				operand := operation[3]

				left := operation[2]
				right := operation[4]

				var (
					leftErr  bool
					rightErr bool
				)

				leftNum, err := strconv.Atoi(left)
				if err != nil {
					leftErr = true
				}
				rightNum, err := strconv.Atoi(right)
				if err != nil {
					rightErr = true
				}
				var leftValue int
				var rightValue int
				if rightErr {
					rightValue = item
				} else {
					rightValue = rightNum
				}
				if leftErr {
					leftValue = item
				} else {
					leftValue = leftNum
				}

				worryLevel := doOperation(operand, leftValue, rightValue)
				worryLevel = getWorryLevel(worryLevel)
				m.inspectCount++
				if worryLevel%m.divisibleBy == 0 {
					ifTrueMonkey := &monkeys[m.ifTrue]
					ifTrueMonkey.items = append(ifTrueMonkey.items, worryLevel)
				} else {
					ifFalseMonkey := &monkeys[m.ifFalse]
					ifFalseMonkey.items = append(ifFalseMonkey.items, worryLevel)
				}
				m.items = m.items[1:]

			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})
	fmt.Println(monkeys[0].inspectCount * monkeys[1].inspectCount)

}
