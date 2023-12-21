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

type cubeSet struct {
	red   int
	blue  int
	green int
}

func main() {
	inputPath := flag.String("path", "./input.txt", "Enter file path")
	maxRed := flag.Int("red", 0, "Enter total count of red cubes in bag")
	maxGreen := flag.Int("green", 0, "Enter total count of green cubes in bag")
	maxBlue := flag.Int("blue", 0, "Enter total count of blue cubes in bag")

	flag.Parse()

	if *maxBlue == 0 && *maxGreen == 0 && *maxRed == 0 {
		log.Fatal("Max count of each colored cube is required")
	}

	file, err := os.Open(*inputPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	total := 0
	totalOfMinimumCubeCount := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineArr := strings.Split(line, ":")
		game := lineArr[0]
		gameIDstr := strings.Split(game, " ")[1]
		gameID, _ := strconv.Atoi(gameIDstr)
		cubes := lineArr[1]
		cubesArr := strings.Split(cubes, ";")
		isPossible := true

		minimumCubeCount := cubeSet{
			red:   0,
			blue:  0,
			green: 0,
		}

		for _, cubeSetStr := range cubesArr {
			cubeSetArr := strings.Split(cubeSetStr, ",")
			for _, cubeStr := range cubeSetArr {
				cube := strings.Split(strings.TrimSpace(cubeStr), " ")
				cubeCountStr := cube[0]
				color := cube[1]
				cubeCount, _ := strconv.Atoi(cubeCountStr)
				switch color {
				case "red":
					if cubeCount > *maxRed {
						isPossible = false
					}
					if cubeCount > minimumCubeCount.red {
						minimumCubeCount.red = cubeCount
					}
					break
				case "green":
					if cubeCount > *maxGreen {
						isPossible = false
					}
					if cubeCount > minimumCubeCount.green {
						minimumCubeCount.green = cubeCount
					}
					break
				case "blue":
					if cubeCount > *maxBlue {
						isPossible = false
					}
					if cubeCount > minimumCubeCount.blue {
						minimumCubeCount.blue = cubeCount
					}
					break
				default:
					isPossible = false

				}

			}
		}
		if isPossible {
			total += gameID
		}
		totalOfMinimumCubeCount += (minimumCubeCount.blue * minimumCubeCount.red * minimumCubeCount.green)

	}
	fmt.Printf("What is the sum of the IDs of those games?: %d\nWhat is the sum of the power of these sets?: %d\n", total, totalOfMinimumCubeCount)

}
