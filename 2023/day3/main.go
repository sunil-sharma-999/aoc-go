package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputPath := flag.String("input", "./input.txt", "Enter input file path")
	flag.Parse()

	file, err := os.Open(*inputPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println(line)
	}

}
