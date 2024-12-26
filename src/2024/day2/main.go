package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Dir int

const (
	IDLE Dir = iota
	INCREASING
	DECREASING
)

func IsSafe(row []int) bool {
	rowLen := len(row)

	dir := IDLE

	for i := 0; i < rowLen-1; i++ {
		diff := row[i] - row[i+1]

		if i != 0 && dir == IDLE {
			return false
		} else {
			if dir == IDLE {
				if diff > 0 {
					dir = INCREASING
				} else if diff < 0 {
					dir = DECREASING
				} else {
					return false
				}
			} else if diff < 0 && dir == INCREASING {
				return false
			} else if diff > 0 && dir == DECREASING {
				return false
			}
		}

		abs := int(math.Abs(float64(diff)))
		if abs > 3 || abs < 1 {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rows := [][]int{}
	whitespaceRegex := regexp.MustCompile(`\s+`)

	for scanner.Scan() {
		rowNumStr := whitespaceRegex.Split(scanner.Text(), -1)
		rowNums := []int{}
		for _, numStr := range rowNumStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			rowNums = append(rowNums, num)
		}
		rows = append(rows, rowNums)
	}

	if scanner.Err() != nil {
		panic("error reading standard input")
	}

	safeReportsPart1 := 0
	safeReportsPart2 := 0

ROW_LOOP:
	for _, row := range rows {
		rowLen := len(row)
		if IsSafe(row) {
			safeReportsPart1++
			safeReportsPart2++
		} else {
			for i := 0; i < rowLen; i++ {
				slice := []int{}
				slice = append(slice, row[:i]...)
				slice = append(slice, row[i+1:]...)
				if IsSafe(slice) {
					safeReportsPart2++
					continue ROW_LOOP
				}
			}
		}
	}

	fmt.Printf("\nPart 1: How many reports are safe?\n%d\n", safeReportsPart1)
	fmt.Printf("Part 2: How many reports are now safe?\n%d\n", safeReportsPart2)
}
