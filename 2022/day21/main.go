package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	name      string
	operation string
}

func MakeMonkey(s string) (monkey Monkey) {
	parts := strings.Split(s, ": ")
	monkey.name = parts[0]
	monkey.operation = parts[1]
	return monkey
}

func MonkeyMath(monkeys List[Monkey], monkey *Monkey) Polynomial {
	if number, err := strconv.ParseInt(monkey.operation, 10, 0); err == nil {
		return Polynomial{ListComparable: ListComparable[float64]{List: []float64{float64(number), 0}}}
	} else if monkey.operation == "x" {
		return Polynomial{ListComparable: ListComparable[float64]{List: []float64{0, 1}}}
	}

	parts := strings.Split(monkey.operation, " ")
	polynomial1 := MonkeyMath(monkeys, monkeys.Find(func(monkey Monkey) bool { return monkey.name == parts[0] }))
	polynomial2 := MonkeyMath(monkeys, monkeys.Find(func(monkey Monkey) bool { return monkey.name == parts[2] }))
	switch parts[1] {
	case "+":
		return polynomial1.Add(polynomial2)
	case "-":
		return polynomial1.Subtract(polynomial2)
	case "*":
		return polynomial1.Multiply(polynomial2)
	case "/":
		return polynomial1.Divide(polynomial2.List[0])
	default:
		panic("operator " + parts[1] + " not implemented")
	}
}

func step1(input string) int {
	monkeys := List[Monkey](operators.Map(strings.Split(input, "\n"), MakeMonkey))
	rootIndex := monkeys.FindIndex(func(monkey Monkey) bool {
		return monkey.name == "root"
	})

	polynomial := MonkeyMath(monkeys, &monkeys[rootIndex])
	return int(polynomial.List[0])
}

func step2(input string) int {
	monkeys := List[Monkey](operators.Map(strings.Split(input, "\n"), MakeMonkey))
	monkeys[monkeys.FindIndex(func(monkey Monkey) bool {
		return monkey.name == "humn"
	})].operation = "x"

	rootIndex := monkeys.FindIndex(func(monkey Monkey) bool {
		return monkey.name == "root"
	})
	parts := strings.Split(monkeys[rootIndex].operation, " ")
	parts[1] = "-"
	monkeys[rootIndex].operation = strings.Join(parts, " ")

	polynomial := MonkeyMath(monkeys, &monkeys[rootIndex])
	return -int(polynomial.List[0] / polynomial.List[1])
}

func main() {
	const title, day = "--- Day 21: Monkey Math ---", "2022/day21/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 152, "example step1")
	utils.AssertEqual(step2(example), 301, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 31017034894002, "step1")
	utils.AssertEqual(step2(input), 3555057453229, "step2")
}
