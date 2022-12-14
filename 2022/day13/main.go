package main

import (
	"adventofcode/models"
	"adventofcode/utils"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func isInt(element any) bool {
	return reflect.TypeOf(element).Kind() == reflect.Float64
}

func isList(element any) bool {
	return reflect.TypeOf(element).Kind() == reflect.Slice
}

func compare(list1 models.Stack[any], list2 models.Stack[any]) bool {
	if len(list1) == 0 {
		return true
	}
	if len(list2) == 0 {
		return false
	}
	if isInt(list1[0]) && isInt(list2[0]) {
		element1, _ := list1.Pop()
		element2, _ := list2.Pop()
		fmt.Println(element1, element2)
		if element1.(float64) == element2.(float64) {
			return compare(list1, list2)
		} else {
			return element1.(float64) < element2.(float64)
		}
	} else if isList(list1[0]) && isList(list2[0]) {
		fmt.Println("list")
	} else {
		fmt.Println("default")
	}
	return false
}

func step1(input string) (result int) {
	var pairs models.Stack[string] = strings.Split(input, "\n\n")
	for index, pair := range pairs {
		parts := strings.Split(pair, "\n")
		var list1, list2 models.Stack[any]
		_ = json.Unmarshal([]byte(parts[0]), &list1)
		_ = json.Unmarshal([]byte(parts[1]), &list2)

		if compare(list1, list2) {
			result += index + 1
		}
	}
	return result
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 13: Distress Signal ---", "2022/day13/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), -1, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")

	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
