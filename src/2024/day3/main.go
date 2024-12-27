package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var matcher = regexp.MustCompile(`(don\'t\(\))|(do\(\))|(mul\(\d+\,\d+\))`)
var digitMatcher = regexp.MustCompile(`\d+`)

func Mul(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mul := nums[0]
	for i := 1; i < len(nums); i++ {
		mul = mul * nums[i]
	}
	return mul
}

func ParseMul(str string) []int {
	numsStr := digitMatcher.FindAllString(str, -1)
	nums := make([]int, len(numsStr))
	for i, numStr := range numsStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

type Executer struct {
	Input        string
	Enabled      bool
	Pos          int
	Sum          int
	CheckEnabled bool
}

func (e *Executer) ExecuteNext() int {
	if e.Pos >= len(e.Input) {
		return 0
	}
	index := matcher.FindStringIndex(e.Input[e.Pos:])

	if len(index) == 0 {
		return 0
	}

	str := e.Input[e.Pos+index[0] : e.Pos+index[1]]
	e.Pos += index[1]
	if str == "don't()" {
		e.Enabled = false
	} else if str == "do()" {
		e.Enabled = true
	} else if str[:4] == "mul(" {
		nums := ParseMul(str)
		if e.Enabled || !e.CheckEnabled {
			e.Sum += Mul(nums)
		}
	}
	return 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines string
	for scanner.Scan() {
		lines += scanner.Text()
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	executer1 := Executer{Input: lines, Enabled: true, Pos: 0, Sum: 0, CheckEnabled: false}
	executer2 := Executer{Input: lines, Enabled: true, Pos: 0, Sum: 0, CheckEnabled: true}

	for {
		shouldContinue := executer1.ExecuteNext()
		if shouldContinue == 0 {
			break
		}
	}

	for {
		shouldContinue := executer2.ExecuteNext()
		if shouldContinue == 0 {
			break
		}
	}

	fmt.Printf("\nPart 1: What do you get if you add up all of the results of the multiplications?\n%d\n", executer1.Sum)
	fmt.Printf("\nPart 2: What do you get if you add up all of the results of just the enabled multiplications?\n%d\n", executer2.Sum)
}
