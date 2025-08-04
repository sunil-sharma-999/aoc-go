package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findIndex(numToMatch int, arr []int) int {
	for i, num := range arr {
		if num == numToMatch {
			return i
		}
	}
	return -1
}

func checkOrder(nums []int, rules map[int][]int) bool {
	for numIndex, num := range nums {
		// assume order is correct if there is no order entry of that number
		if rules[num] == nil {
			continue
		}
		// loop over order values of pageNum where pageNum should be first
		for _, orderNum := range rules[num] {
			orderNumIndex := findIndex(orderNum, nums)
			if orderNumIndex < numIndex && orderNumIndex != -1 {
				return false
			}
		}
	}
	return true
}

func orderCorrection(nums []int, rules map[int][]int) []int {
	for !checkOrder(nums, rules) {
		for numIndex := 0; numIndex < len(nums); numIndex++ {
			num1 := nums[numIndex]
			if numIndex == 0 {
				continue
			}
			currNumRuleOrderNums := rules[num1]

			if currNumRuleOrderNums == nil {
				continue
			}
			for _, orderNum := range currNumRuleOrderNums {
				foundOrderNumIndex := findIndex(orderNum, nums)
				if foundOrderNumIndex == -1 || numIndex < foundOrderNumIndex {
					continue
				}
				nums[numIndex] = orderNum
				nums[foundOrderNumIndex] = num1
				numIndex = foundOrderNumIndex
			}
		}
	}
	return nums
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
	invalidPageNumRows := [][]int{}

	// loop over all rows
	for _, pageNumberRow := range pageNumberRows {
		isValid := checkOrder(pageNumberRow, orders)
		if isValid {
			validPageNumRows = append(validPageNumRows, pageNumberRow)
		} else {
			invalidPageNumRows = append(invalidPageNumRows, pageNumberRow)
		}
	}

	sum := 0

	for _, validRow := range validPageNumRows {
		sum += validRow[len(validRow)/2]
	}

	fmt.Println("What do you get if you add up the middle page number from those correctly-ordered updates?", sum)

	// 2nd Part
	sum = 0

	correctedOrders := [][]int{}

	for _, invalidRow := range invalidPageNumRows {
		correctedOrders = append(correctedOrders, orderCorrection(invalidRow, orders))
	}

	for _, correctedOrder := range correctedOrders {
		sum += correctedOrder[len(correctedOrder)/2]
	}

	fmt.Println("What do you get if you add up the middle page numbers after correctly ordering just those updates?", sum)
}
