package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func formatSeeds(str string) []int {
	arr := strings.Fields(strings.Split(strings.TrimSpace(str), ":")[1])
	numArr := []int{}
	for _, numStr := range arr {
		numArr = append(numArr, ToNum(numStr))
	}
	return numArr
}

func findMin(arr []int) int {
	min := 0
	if len(arr) >= 1 {
		min = arr[0]
	}
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	return min
}

func formatMaps(str string) [][]int {
	dataStr := strings.Split(str, ":")[1]
	arrOfArrStr := strings.Split(strings.TrimSpace(dataStr), "\n")
	arrOfArrNum := [][]int{}
	for i, arrStr := range arrOfArrStr {
		arrOfArrNum = append(arrOfArrNum, []int{})
		for _, numStr := range strings.Fields(arrStr) {
			arrOfArrNum[i] = append(arrOfArrNum[i], ToNum(numStr))
		}
	}
	return arrOfArrNum
}

func findCorrespond(num int, arr [][]int) int {
	mappedNum := num

loop:
	for _, details := range arr {
		destination := details[0]
		source := details[1]
		rangeLength := details[2]
		if num >= source && num < source+rangeLength {
			mappedNum = (destination - source) + num
			break loop
		}
	}
	return mappedNum
}

func main() {
	inputPath := flag.String("input", "./input.txt", "Enter input file path")
	flag.Parse()

	file, err := os.Open(*inputPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	dataArr := []string{}
	str := ""

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			dataArr = append(dataArr, strings.TrimSpace(str))
			str = ""
		} else {
			str += line + "\n"
		}
	}
	dataArr = append(dataArr, strings.TrimSpace(str))
	// cleanup str
	str = ""

	seeds := formatSeeds(dataArr[0])

	locations := []int{}
	for _, seed := range seeds {
		location := seed
		for _, mapper := range dataArr[1:] {
			location = findCorrespond(location, formatMaps(mapper))
		}
		locations = append(locations, location)
	}
	fmt.Printf("\nWhat is the lowest location number that corresponds to any of the initial seed numbers? %d\n\n", findMin(locations))
}
