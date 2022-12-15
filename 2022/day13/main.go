package main

import (
	"adventofcode/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type SortablePackets [][]any

func (packets SortablePackets) Len() int {
	return len(packets)
}

func (packets SortablePackets) Less(i, j int) bool {
	return compare(packets[i], packets[j]) == Inferior
}

func (packets SortablePackets) Swap(i, j int) {
	packets[i], packets[j] = packets[j], packets[i]
}

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
	var comparison Comparison
	if isInt(element1) && isInt(element2) {
		if element1.(float64) == element2.(float64) {
			comparison = Equal
		} else if element1.(float64) < element2.(float64) {
			comparison = Inferior
		} else {
			comparison = Superior
		}
	} else if isList(element1) && isList(element2) {
		comparison = compare(element1.([]any), element2.([]any))
	} else if isInt(element1) && isList(element2) {
		comparison = compare([]any{element1}, element2.([]any))
	} else if isList(element1) && isInt(element2) {
		comparison = compare(element1.([]any), []any{element2})
	} else {
		panic("not implemented")
	}

	if comparison == Equal {
		return compare(list1, list2)
	} else {
		return comparison
	}
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

func step2(input string) int {
	var packets [][]any
	firstPacket, secondPacket := []any{[]any{2.0}}, []any{[]any{6.0}}
	packets = append(packets, firstPacket, secondPacket)
	for _, pair := range strings.Split(input, "\n\n") {
		parts := strings.Split(pair, "\n")
		var list1, list2 []any
		_ = json.Unmarshal([]byte(parts[0]), &list1)
		_ = json.Unmarshal([]byte(parts[1]), &list2)

		packets = append(packets, list1, list2)
	}
	sort.Sort(SortablePackets(packets))
	firstPacketIndex := operators.FindIndex(packets, func(element []any) bool {
		return compare(element, firstPacket) == Equal
	})
	secondPacketIndex := operators.FindIndex(packets, func(element []any) bool {
		return compare(element, secondPacket) == Equal
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
