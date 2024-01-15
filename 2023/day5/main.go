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

type Data struct {
	seeds                 []int
	seedsToSoil           map[int]int
	soilToFertilizer      map[int]int
	fertilizerToWater     map[int]int
	waterToLight          map[int]int
	lightToTemperature    map[int]int
	temperatureToHumidity map[int]int
	humidityToLocation    map[int]int
}

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

func mapSeedsToResource(arr [][]int) map[int]int {
	numsMap := map[int]int{}
	for i := 0; i < 100; i++ {
		numsMap[i] = i
	}
	for _, details := range arr {
		destination := details[0]
		source := details[1]
		rangeLength := details[2]
		for i := source; i < source+rangeLength; i++ {
			numsMap[i] = destination + (i - source)
		}
	}
	return numsMap
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
	// cleanup
	str = ""

	data := Data{
		seeds:                 formatSeeds(dataArr[0]),
		seedsToSoil:           mapSeedsToResource(formatMaps(dataArr[1])),
		soilToFertilizer:      mapSeedsToResource(formatMaps(dataArr[2])),
		fertilizerToWater:     mapSeedsToResource(formatMaps(dataArr[3])),
		waterToLight:          mapSeedsToResource(formatMaps(dataArr[4])),
		lightToTemperature:    mapSeedsToResource(formatMaps(dataArr[5])),
		temperatureToHumidity: mapSeedsToResource(formatMaps(dataArr[6])),
		humidityToLocation:    mapSeedsToResource(formatMaps(dataArr[7])),
	}
	locations := []int{}
	for _, s := range data.seeds {
		locations = append(locations, data.humidityToLocation[data.temperatureToHumidity[data.lightToTemperature[data.waterToLight[data.fertilizerToWater[data.soilToFertilizer[data.seedsToSoil[s]]]]]]])

	}
	fmt.Printf("\nWhat is the lowest location number that corresponds to any of the initial seed numbers? %d\n\n", findMin(locations))
}
