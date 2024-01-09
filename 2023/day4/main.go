package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func sum(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total

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
	re := regexp.MustCompile(`\:|\|`)
	total := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineArr := re.Split(line, -1)
		winningNumbers := strings.Fields(strings.TrimSpace(lineArr[1]))
		numbers := strings.Fields(strings.TrimSpace(lineArr[2]))

		matches := []int{}

		for _, numChar := range numbers {
			alreadyMatches := map[string]bool{}
		loop:
			for _, winningNumChar := range winningNumbers {
				_, ok := alreadyMatches[numChar]
				if strings.TrimSpace(numChar) == strings.TrimSpace(winningNumChar) && (!ok) {
					alreadyMatches[strings.TrimSpace(numChar)] = true
					if len(matches) < 2 {
						matches = append(matches, 1)
					} else {
						matches = append(matches, matches[len(matches)-1:][0]*2)
					}
					break loop
				}
			}
		}
		total += sum(matches)
	}
	fmt.Printf("How many points are they worth in total? %d\n", total)
}
