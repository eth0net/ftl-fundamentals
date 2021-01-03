// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Divide takes two numbers and returns the result of dividing the first from
// the second.
func Divide(in ...float64) (out float64, err error) {
	if len(in) < 2 {
		return 0, fmt.Errorf("too few values %d, need at least 2", len(in))
	}

	out = in[0]
	for _, n := range in[1:] {
		if n == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		out /= n
	}
	return out, nil
}

// Multiply takes two numbers and returns the result of multiplying them
// together.
func Multiply(a, b float64) float64 {
	return a * b
}

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Sqrt takes one number and returns its square root.
// An error will be returned for negative values.
func Sqrt(in float64) (out float64, err error) {
	if in < 0 {
		return 0, fmt.Errorf("cannot square root a negative number: %f", in)
	}
	return math.Sqrt(in), nil
}
