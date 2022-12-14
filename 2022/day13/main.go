package main

import (
	"adventofcode/models"
	"adventofcode/utils"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Comparison string

const (
	Inferior Comparison = "Inferior"
	Superior            = "Superior"
	Equal               = "Equal"
)

func isInt(element any) bool {
	return reflect.TypeOf(element).Kind() == reflect.Float64
}

func isList(element any) bool {
	return reflect.TypeOf(element).Kind() == reflect.Slice
}

func compare(list1 models.Stack[any], list2 models.Stack[any]) Comparison {
	if len(list1) == 0 && len(list2) == 0 {
		return Equal
	} else if len(list1) == 0 {
		return Inferior
	} else if len(list2) == 0 {
		return Superior
	}

	element1, _ := list1.Shift()
	element2, _ := list2.Shift()
	if isInt(element1) && isInt(element2) {
		if element1.(float64) == element2.(float64) {
			return compare(list1, list2)
		} else if element1.(float64) < element2.(float64) {
			return Inferior
		} else {
			return Superior
		}
	} else if isList(element1) && isList(element2) {
		comparison := compare(element1.([]any), element2.([]any))
		if comparison == Equal {
			return compare(list1, list2)
		} else {
			return comparison
		}
	} else if isInt(element1) && isList(element2) {
		comparison := compare([]any{element1}, element2.([]any))
		if comparison == Equal {
			return compare(list1, list2)
		} else {
			return comparison
		}
	} else if isList(element1) && isInt(element2) {
		comparison := compare(element1.([]any), []any{element2})
		if comparison == Equal {
			return compare(list1, list2)
		} else {
			return comparison
		}
	} else {
		panic("not defined")
	}
	return Equal
}

func step1(input string) (result int) {
	var pairs models.Stack[string] = strings.Split(input, "\n\n")
	for index, pair := range pairs {
		parts := strings.Split(pair, "\n")
		var list1, list2 models.Stack[any]
		_ = json.Unmarshal([]byte(parts[0]), &list1)
		_ = json.Unmarshal([]byte(parts[1]), &list2)

		if compare(list1, list2) == Inferior {
			result += index + 1
		}
	}
	return result
}

func BubbleSort(packets *models.Stack[models.Stack[any]]) {
	isDone := false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(*packets)-1 {
			if compare((*packets)[i], (*packets)[i+1]) == Superior {
				(*packets)[i], (*packets)[i+1] = (*packets)[i+1], (*packets)[i]
				isDone = false
			}
			i++
		}
	}
}

func step2(input string) int {
	var packets models.Stack[models.Stack[any]]
	packets = append(packets, []any{[]any{2.0}}, []any{[]any{6.0}})
	var pairs models.Stack[string] = strings.Split(input, "\n\n")
	for _, pair := range pairs {
		parts := strings.Split(pair, "\n")
		var list1, list2 models.Stack[any]
		_ = json.Unmarshal([]byte(parts[0]), &list1)
		_ = json.Unmarshal([]byte(parts[1]), &list2)

		packets = append(packets, list1, list2)
	}
	BubbleSort(&packets)
	firstPacketIndex := packets.Find(func(element models.Stack[any]) bool {
		return len(element) == 1 && isList(element[0]) && len(element[0].([]any)) == 1 && element[0].([]any)[0] == 2.0
	})
	secondPacketIndex := packets.Find(func(element models.Stack[any]) bool {
		return len(element) == 1 && isList(element[0]) && len(element[0].([]any)) == 1 && element[0].([]any)[0] == 6.0
	})
	return (firstPacketIndex + 1) * (secondPacketIndex + 1)
}

func main() {
	const title, day = "--- Day 13: Distress Signal ---", "2022/day13/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 13, "example step1")
	utils.AssertEqual(step2(example), 140, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 6369, "step1")
	utils.AssertEqual(step2(input), 25800, "step2")
}
