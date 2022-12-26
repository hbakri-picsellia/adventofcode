package main

import (
	"adventofcode/mathInt"
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func runeToInt(r rune) int {
	switch r {
	case '2', '1', '0':
		return int(r - '0')
	case '-':
		return -1
	case '=':
		return -2
	default:
		panic("rune undefined")
	}
}

func intToRune(i int) rune {
	switch i {
	case 2, 1, 0:
		return rune(strconv.FormatInt(int64(i), 10)[0])
	case -1:
		return '-'
	case -2:
		return '='
	default:
		panic("int undefined")
	}
}

type SNAFU struct{}

func (snafu SNAFU) Decode(s string) (result int) {
	for index, value := range s {
		result += runeToInt(value) * mathInt.Pow(5, len(s)-1-index)
	}
	return
}

func (snafu SNAFU) Encode(i int) string {
	var digits List[int] = operators.Map([]rune(strconv.FormatInt(int64(i), 5)), func(r rune) int {
		return utils.ParseInt(string(r))
	})
	digits.Insert(0, 0)
	for j := len(digits) - 1; j > 0; j-- {
		if digits[j] > 2 {
			digits[j] -= 5
			digits[j-1] += 1
		}
	}
	return strings.TrimLeft(string(operators.Map(digits, intToRune)), "0")
}

func step1(input string) string {
	return SNAFU{}.Encode(operators.Sum(operators.Map(strings.Split(input, "\n"), SNAFU{}.Decode)))
}

func step2(input string) string {
	return ""
}

func main() {
	const title, day = "--- Day 25: Full of Hot Air ---", "2022/day25/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), "2=-1=0", "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), "2-=2-0=-0-=0200=--21", "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
