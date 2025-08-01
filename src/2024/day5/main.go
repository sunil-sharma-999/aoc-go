package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isArrayEqual(array1, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i := range array1 {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}

func findIndex(numToMatch int, arr []int) int {
	for i, num := range arr {
		if num == numToMatch {
			return i
		}
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	orders := map[int][]int{}
	pageNumberRows := [][]int{}

scannerLoop:
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			chars := strings.Split(line, "|")[0:2]
			keyNum, err := strconv.Atoi(chars[0])
			if err != nil {
				fmt.Println("Expected number got: ", chars[0], "\n", err)
				continue scannerLoop
			}
			num, err := strconv.Atoi(chars[1])
			if err != nil {
				fmt.Println("Expected number got: ", chars[1], "\n", err)
				continue scannerLoop
			}
			if orders[keyNum] == nil {
				orders[keyNum] = []int{}
			}
			orders[keyNum] = append(orders[keyNum], num)
		} else if strings.Contains(line, ",") {
			numsStr := strings.Split(line, ",")
			nums := []int{}
			for _, numStr := range numsStr {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					fmt.Println("Expected number got: ", numStr, "\n", err)
					continue scannerLoop
				}
				nums = append(nums, num)
			}
			pageNumberRows = append(pageNumberRows, nums)
		}
	}

	validPageNumRows := [][]int{}

	// loop over all rows
rowsLoop:
	for _, pageNumberRow := range pageNumberRows {
		// loop over numbers of row
		isValid := true
		for pageNumIndex, pageNum := range pageNumberRow {
			// assume order is correct if there is no order entry of that number
			if orders[pageNum] == nil {
				continue
			}
			// loop over order values of pageNum where pageNum should be first
			for _, orderNum := range orders[pageNum] {
				orderNumIndex := findIndex(orderNum, pageNumberRow)
				if orderNumIndex < pageNumIndex && orderNumIndex != -1 {
					isValid = false
					continue rowsLoop
				}
			}
		}
		if isValid {
			validPageNumRows = append(validPageNumRows, pageNumberRow)
		}
	}

	sum := 0

	for _, validRow := range validPageNumRows {
		sum += validRow[len(validRow)/2]
	}

	fmt.Println("What do you get if you add up the middle page number from those correctly-ordered updates?", sum)
}
