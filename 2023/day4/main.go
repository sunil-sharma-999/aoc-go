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

func countScratchcards(cardNum int, mapper map[int]int, cardMatcher map[int][]int) {
	matches := cardMatcher[cardNum]
	mapper[cardNum]++
	for _, nextCardNum := range matches {
		countScratchcards(nextCardNum, mapper, cardMatcher)
	}
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
	mapper := map[int]int{}
	allCards := [][][]string{}
	cardMatcher := map[int][]int{}

	cardIndex := -1
	for scanner.Scan() {
		cardIndex++
		mapper[cardIndex+1] = 0

		line := strings.TrimSpace(scanner.Text())
		lineArr := re.Split(line, -1)
		winningNumbers := strings.Fields(strings.TrimSpace(lineArr[1]))
		numbers := strings.Fields(strings.TrimSpace(lineArr[2]))
		matches := []int{}
		allCards = append(allCards, [][]string{numbers, winningNumbers})

		matchingNumbers := 0

		for _, numChar := range numbers {
		loop:
			for _, winningNumChar := range winningNumbers {
				if strings.TrimSpace(numChar) == strings.TrimSpace(winningNumChar) {
					if len(matches) < 2 {
						matches = append(matches, 1)
					} else {
						matches = append(matches, matches[len(matches)-1:][0]*2)
					}
					// for part 2
					matchingNumbers++
					break loop
				}
			}
		}
		total += sum(matches)

		// winning cards
		for i := 0; i < matchingNumbers; i++ {
			cardMatcher[cardIndex+1] = append(cardMatcher[cardIndex+1], cardIndex+2+i)
		}

	}
	// count winning cards
	for cardNum := 1; cardNum <= len(allCards); cardNum++ {
		countScratchcards(cardNum, mapper, cardMatcher)
	}
	totalScratchCards := 0

	for _, count := range mapper {
		totalScratchCards += count
	}

	fmt.Printf("How many points are they worth in total? %d\n", total)
	fmt.Printf("how many total scratchcards do you end up with? %d\n", totalScratchCards)

}
