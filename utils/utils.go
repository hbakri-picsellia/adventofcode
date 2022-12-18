package utils

import (
	"adventofcode/operators"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseFileToString(filename string) string {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(fileContent), "\n")
}

func ParseStringToInt(s string) int {
	num, _ := strconv.ParseInt(s, 10, 0)
	return int(num)
}

func ParseStringToIntList(text string, separator string) []int {
	return operators.Map(strings.Split(text, separator), ParseStringToInt)
}

func AssertEqual[T comparable](input1 T, input2 T, message string) {
	if input1 == input2 {
		fmt.Println(message, "succeeded !")
	} else {
		fmt.Println(message, "failed...", "output:", input1, "expected:", input2)
	}
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Remove[T any](s []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func GetStringBetween(s string, sep1 string, sep2 string) string {
	return strings.TrimSpace(strings.SplitN(strings.SplitAfterN(s, sep1, 2)[1], sep2, 2)[0])
}
