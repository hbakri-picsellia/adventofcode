package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func BinaryDiagnostic(input string) []float32 {
	list := strings.Split(input, "\n")
	nbBinaryNumbers := len(list)
	return utils.Map(
		utils.Reduce(list, func(reduced []int, value string) []int {
			binaryNumber := utils.ParseStringToIntList(value, "")
			for index, _ := range reduced {
				reduced[index] += binaryNumber[index]
			}
			return reduced
		}, make([]int, len(list[0]))), func(nbApparitions int) float32 {
			return float32(nbApparitions) / float32(nbBinaryNumbers)
		})
}

func GammaRate(binaryDiagnostic []float32) string {
	return strings.Join(utils.Map(binaryDiagnostic, func(frequency float32) string {
		result := 0
		if frequency > 0.5 {
			result = 1
		}
		return strconv.FormatInt(int64(result), 2)
	}), "")
}

func EpsilonRate(binaryDiagnostic []float32) string {
	return strings.Join(utils.Map(binaryDiagnostic, func(frequency float32) string {
		result := 0
		if frequency < 0.5 {
			result = 1
		}
		return strconv.FormatInt(int64(result), 2)
	}), "")
}

func step1(input string) int {
	gamma, _ := strconv.ParseInt(GammaRate(BinaryDiagnostic(input)), 2, 0)
	epsilon, _ := strconv.ParseInt(EpsilonRate(BinaryDiagnostic(input)), 2, 0)
	return int(gamma * epsilon)
}

func step2(input string) int {
	return 0
}

func main() {
	fmt.Println("--- Day 3: Binary Diagnostic ---")

	example := utils.ParseFileToString("2021/day03/example.txt")
	utils.AssertEqual(step1(example), 198, "example step1")
	//utils.AssertEqual(step2(example), 900, "example step2")

	input := utils.ParseFileToString("2021/day03/input.txt")
	utils.AssertEqual(step1(input), 1648020, "step1")
	//utils.AssertEqual(step2(input), 1759818555, "step2")
}
