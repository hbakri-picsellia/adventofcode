package structs

import (
	"fmt"
	"math"
	"strconv"
)

type Polynomial struct {
	ListComparable[float64]
}

func MakePolynomial(n int) Polynomial {
	return Polynomial{ListComparable: ListComparable[float64]{List: make([]float64, n)}}
}

func (polynomial *Polynomial) Add(polynomial2 Polynomial) Polynomial {
	result := MakePolynomial(int(math.Max(float64(polynomial.Len()), float64(polynomial2.Len()))))
	for index := range result.List {
		if index < polynomial.Len() {
			result.List[index] += polynomial.List[index]
		}
		if index < polynomial2.Len() {
			result.List[index] += polynomial2.List[index]
		}
	}
	return result
}

func (polynomial *Polynomial) Subtract(polynomial2 Polynomial) Polynomial {
	result := MakePolynomial(int(math.Max(float64(polynomial.Len()), float64(polynomial2.Len()))))
	for index := range result.List {
		if index < polynomial.Len() {
			result.List[index] += polynomial.List[index]
		}
		if index < polynomial2.Len() {
			result.List[index] -= polynomial2.List[index]
		}
	}
	return result
}

func (polynomial *Polynomial) Multiply(polynomial2 Polynomial) Polynomial {
	n, m := polynomial.Len(), polynomial2.Len()
	result := MakePolynomial(n + m - 1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result.List[i+j] += polynomial.List[i] * polynomial2.List[j]
		}
	}
	return result
}

func (polynomial *Polynomial) Divide(constant float64) Polynomial {
	result := MakePolynomial(polynomial.Len())
	for index := range result.List {
		result.List[index] = polynomial.List[index] / constant
	}
	return result
}

func (polynomial *Polynomial) Display() {
	for index := range polynomial.List {
		fmt.Print(polynomial.List[index])
		if index > 0 {
			fmt.Print("x^" + strconv.FormatInt(int64(index), 10))
		}
		if index < polynomial.Len()-1 {
			fmt.Print(" + ")
		}
	}
}
