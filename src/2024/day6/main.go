package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	UP    = '^'
	DOWN  = 'v'
	LEFT  = '<'
	RIGHT = '>'
)

func updateVisited(x int, y int, visited *map[string]bool) {
	key := fmt.Sprintf("%d-%d", x, y)
	if !(*visited)[key] {
		(*visited)[key] = true
	}
}

func getIsBlocked(lab *[][]rune, x, y int) bool {
	if y >= len(*lab) || y < 0 {
		return true
	}
	row := (*lab)[y]

	if x >= len(row) || x < 0 {
		return true
	}
	if row[x] == '#' {
		return true
	} else {
		return false
	}

}

func updateLab(lab *[][]rune, oldCurr, newCurr [2]int, dir rune) {
	x, y := newCurr[0], newCurr[1]
	oldX, oldY := oldCurr[0], oldCurr[1]
	(*lab)[oldY][oldX] = '.'
	(*lab)[y][x] = dir
}

func moveForward(curr *[2]int, lab *[][]rune, visited *map[string]bool) bool {
	x, y := (*curr)[0], (*curr)[1]
	dir := (*lab)[y][x]

	var dx, dy int
	newDir := dir

	switch dir {
	case UP:
		dx, dy = 0, -1
		newDir = RIGHT

	case DOWN:
		dx, dy = 0, 1
		newDir = LEFT
	case LEFT:
		dx, dy = -1, 0
		newDir = UP

	case RIGHT:
		dx, dy = 1, 0
		newDir = DOWN
	default:
		panic("Invalid direction")
	}

	//Move until blocked
	for !getIsBlocked(lab, x+dx, y+dy) {
		x += dx
		y += dy
		updateLab(lab, [2]int{x - dx, y - dy}, [2]int{x, y}, dir)
		updateVisited(x, y, visited)
	}

	updateLab(lab, [2]int{x, y}, [2]int{x, y}, newDir)
	(*curr)[0], (*curr)[1] = x, y

	if x <= 0 || y <= 0 || y >= len(*lab)-1 || x >= len((*lab)[0])-1 {
		return true
	}
	return false

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lab := [][]rune{}
	curr := [2]int{0, 0}
	currFound := false

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lab = append(lab, []rune(line))
		if !currFound {
			for x, char := range line {
				if char != '#' && char != '.' {
					currFound = true
					curr = [2]int{x, len(lab) - 1}
				}
			}
		}
	}

	// format of 'x-y': bool
	visited := map[string]bool{}
	// init
	updateVisited(curr[0], curr[1], &visited)
	reachedEdge := false

	for !reachedEdge {
		reachedEdge = moveForward(&curr, &lab, &visited)
	}

	fmt.Println(len(visited))

}
