package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dop251/goja"
)

var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

const SCRIPT = `
	function getString(str) {
		const regex = /(?=([0-9])|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine))/g
		return [...str.matchAll(regex)].map(x => x.filter(Boolean)?.[0]).join("")
	}
`

func main() {
	path := flag.String("path", "./input.txt", "Enter file path")
	flag.Parse()

	file, err := os.Open(*path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// scanner
	scanner := bufio.NewScanner(file)

	total := 0

	// VM
	vm := goja.New()
	_, err = vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}
	getString, ok := goja.AssertFunction(vm.Get("getString"))
	if !ok {
		panic("Not a function")
	}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		res, err := getString(goja.Undefined(), vm.ToValue(line))
		if err != nil {
			panic(err)
		}
		line = res.ToString().String()
		sent := regexp.MustCompile(`(([0-9])|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine))`)

		matches := sent.FindAllStringIndex(line, -1)
		if len(matches) > 0 {
			match1 := matches[0]
			var match2 []int
			if len(matches) == 1 {
				match2 = match1
			} else {
				match2 = matches[len(matches)-1]
			}

			num1Str := line[match1[0]:match1[1]]
			mappedStr, ok := numMap[num1Str]
			if ok {
				num1Str = mappedStr
			}
			num2Str := line[match2[0]:match2[1]]
			mappedStr, ok = numMap[num2Str]
			if ok {
				num2Str = mappedStr
			}
			num, err := strconv.Atoi(num1Str + num2Str)
			if err != nil {
				log.Fatal(err)
			}
			total += num
		}

	}
	fmt.Println(total)
}
