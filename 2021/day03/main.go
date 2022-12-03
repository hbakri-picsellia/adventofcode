package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func rateToDecimal(rate string) int {
	decimal, _ := strconv.ParseInt(rate, 2, 0)
	return int(decimal)
}

func BinaryDiagnostic(binaryNumbers []string) []float64 {
	return operators.Map(
		operators.Reduce(binaryNumbers, func(reduced []int, value string) []int {
			return operators.Add(reduced, utils.ParseStringToIntList(value, ""))
		}, make([]int, len(binaryNumbers[0]))), func(nbApparitions int) float64 {
			return float64(nbApparitions) / float64(len(binaryNumbers))
		})
}

func binaryRate(binaryDiagnostic []float64, f func(float64) float64) string {
	return strings.Join(operators.Map(binaryDiagnostic, func(frequency float64) string {
		return strconv.FormatInt(int64(f(frequency)), 2)
	}), "")
}

func gammaRate(binaryDiagnostic []float64) string {
	return binaryRate(binaryDiagnostic, func(frequency float64) float64 {
		return math.Round(frequency)
	})
}

func epsilonRate(binaryDiagnostic []float64) string {
	return binaryRate(binaryDiagnostic, func(frequency float64) float64 {
		return 1 - math.Round(frequency)
	})
}

func bitCriteria(binaryNumbers []string, binaryRate func([]string) string) string {
	index := 0
	for {
		rateValue := binaryRate(binaryNumbers)
		binaryNumbers = operators.Filter(binaryNumbers, func(binaryNumber string) bool {
			return binaryNumber[index] == rateValue[index]
		})
		if len(binaryNumbers) == 1 {
			return binaryNumbers[0]
		}
		index += 1
	}
}

func oxygenGeneratorRate(binaryNumbers []string) string {
	return bitCriteria(binaryNumbers, func(binaryNumbers []string) string {
		return gammaRate(BinaryDiagnostic(binaryNumbers))
	})
}

func co2ScrubberRate(binaryNumbers []string) string {
	return bitCriteria(binaryNumbers, func(binaryNumbers []string) string {
		return epsilonRate(BinaryDiagnostic(binaryNumbers))
	})
}

func step1(input string) int {
	binaryNumbers := strings.Split(input, "\n")
	binaryDiagnostic := BinaryDiagnostic(binaryNumbers)
	gamma := rateToDecimal(gammaRate(binaryDiagnostic))
	epsilon := rateToDecimal(epsilonRate(binaryDiagnostic))
	return gamma * epsilon
}

func step2(input string) int {
	binaryNumbers := strings.Split(input, "\n")
	oxygenGenerator := rateToDecimal(oxygenGeneratorRate(binaryNumbers))
	co2Scrubber := rateToDecimal(co2ScrubberRate(binaryNumbers))
	return oxygenGenerator * co2Scrubber
}

func main() {
	fmt.Println("--- Day 3: Binary Diagnostic ---")

	example := utils.ParseFileToString("2021/day03/example.txt")
	utils.AssertEqual(step1(example), 198, "example step1")
	utils.AssertEqual(step2(example), 230, "example step2")

	input := utils.ParseFileToString("2021/day03/input.txt")
	utils.AssertEqual(step1(input), 3549854, "step1")
	utils.AssertEqual(step2(input), 3765399, "step2")
}
