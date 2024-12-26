package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	expressionMatcher := regexp.MustCompile(`mul\(\d+\,\d+\)`)
	digitMatcher := regexp.MustCompile(`\d+`)
	expressions := []string{}
	for _, line := range lines {
		expressions = append(expressions, expressionMatcher.FindAllString(line, -1)...)
	}
	sum := 0

	for _, expression := range expressions {
		numsStr := digitMatcher.FindAllString(expression, -1)
		nums := []int{}
		for _, numStr := range numsStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		mul := nums[0]
		for i := 1; i < len(nums); i++ {
			mul = mul * nums[i]
		}
		sum += mul
	}
	fmt.Printf("\nPart 1: What do you get if you add up all of the results of the multiplications?\n%d\n", sum)
}
