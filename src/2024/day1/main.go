package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	inputLines := []string{}

	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if scanner.Err() != nil {
		panic("error reading standard input")
	}

	totalRows := len(inputLines)
	whitespaceRegex := regexp.MustCompile(`\s+`)
	nums1 := []int{}
	nums2 := []int{}

	for i := 0; i < totalRows; i++ {
		lineArr := whitespaceRegex.Split(inputLines[i], -1)
		if len(lineArr) != 2 {
			panic(("Invalid input"))
		}
		num1, err := strconv.Atoi(lineArr[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(lineArr[1])
		if err != nil {
			panic(err)
		}
		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)
	}

	sort.IntSlice(nums1).Sort()
	sort.IntSlice(nums2).Sort()

	sum := 0

	for i := 0; i < len(nums1); i++ {
		abs := math.Abs(float64(nums1[i]) - float64(nums2[i]))
		sum += int(abs)
	}
	fmt.Println()
	fmt.Printf("Part 1: What is the total distance between your lists?\n%d\n", sum)

	sum = 0
	for _, num1 := range nums1 {
		times := 0
		for _, num2 := range nums2 {
			if num1 == num2 {
				times++
			}
		}
		sum += times * num1
	}
	fmt.Println()
	fmt.Printf("Part 2: What is their similarity score?\n%d\n", sum)
}
