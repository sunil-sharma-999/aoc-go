package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x    int
	y    int
	name string
}

func getClose(head *Node, tail *Node) {
	diff := [2]int{head.x - tail.x, head.y - tail.y}
	if math.Abs(float64(diff[0])) > 1 || math.Abs(float64(diff[1])) > 1 {
		if diff[0] > 0 {
			tail.x += 1
		} else if diff[0] < 0 {
			tail.x -= 1
		}
		if diff[1] > 0 {
			tail.y += 1
		} else if diff[1] < 0 {
			tail.y -= 1
		}
	}

}

type Position struct {
	direction string
	steps     int
}

func (p Position) step(n *Node) {
	switch p.direction {
	case "R":
		n.x += 1
	case "U":
		n.y += 1
	case "L":
		n.x += -1
	case "D":
		n.y += -1
	}

}

func getMove(line string) Position {
	strArr := strings.Fields(line)

	steps, err := strconv.Atoi(strArr[1])

	if err != nil {
		panic(err)
	}

	return Position{
		direction: strArr[0],
		steps:     steps,
	}

}

func makeMove(move Position, head *Node, tails *[]*Node, visited *map[string]bool) {
	for i := 0; i < move.steps; i++ {
		if head.name == "head" {
			move.step(head)
		}
		for i, tail := range *tails {
			if i == 0 {
				getClose(head, tail)
			} else {
				getClose((*tails)[i-1], tail)
			}

			if tail.name == "tail" && !(*visited)[fmt.Sprintf("%v-%v", tail.x, tail.y)] {
				(*visited)[fmt.Sprintf("%v,%v", tail.x, tail.y)] = true
			}
		}

	}

}

func main() {
	tailCount := flag.Int("tails", 1, "Enter the number of tails")
	path := flag.String("path", "./2023/day9/input.txt", "Enter file path")
	flag.Parse()

	if *path == "" {
		log.Fatal("Enter file path")
	}

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	visited := map[string]bool{
		"0,0": true,
	}

	head := Node{0, 0, "head"}
	tails := []*Node{}

	for i := 0; i < *tailCount; i++ {
		var name string
		if i == *tailCount-1 {
			name = "tail"
		} else {
			name = fmt.Sprintf("t%v", i+1)
		}
		tails = append(tails, &Node{0, 0, name})
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(
			scanner.Text(), "\n")
		move := getMove(line)
		makeMove(move, &head, &tails, &visited)
	}
	fmt.Println(len(visited))
}
