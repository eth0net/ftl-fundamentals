package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	name        string
	a, b        float64
	want        float64
	errExpected bool
}

func TestDivide(t *testing.T) {
	t.Parallel()

	prefix := "%s: Divide(%f, %f): "

	testCases := []testCase{
		{
			name: "Positive divided by positive is positive",
			a:    1, b: 1, want: 1, errExpected: false,
		},
		{
			name: "Negative divided by negative is positive",
			a:    -1, b: -1, want: 1, errExpected: false,
		},
		{
			name: "Positive divided by negative is negative",
			a:    1, b: -1, want: -1, errExpected: false,
		},
		{
			name: "Negative divided by positive is negative",
			a:    -1, b: 1, want: -1, errExpected: false,
		},
		{
			name: "Positive divided by a fraction is positive",
			a:    2, b: .5, want: 4, errExpected: false,
		},
		{
			name: "Negative divided by a fraction is negative",
			a:    -2, b: .5, want: -4, errExpected: false,
		},
		{
			name: "Positive divided by zero returns an error",
			a:    2, b: 0, want: 0, errExpected: true,
		},
		{
			name: "Negative divided by zero returns an error",
			a:    -2, b: 0, want: 0, errExpected: true,
		},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf(prefix+"Unexpected error status: %t", tc.name, tc.a, tc.b, errReceived)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf(prefix+"want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			name: "Zero times zero is zero",
			a:    0, b: 0, want: 0,
		},
		{
			name: "Positive times positive is positive",
			a:    2, b: 4, want: 8,
		},
		{
			name: "Positive times negative is negative",
			a:    2, b: -4, want: -8,
		},
		{
			name: "Negative times negative is positive",
			a:    -2, b: -4, want: 8,
		},
		{
			name: "Halving a positive",
			a:    0.5, b: 2, want: 1,
		},
		{
			name: "Halving a negative",
			a:    0.5, b: -4, want: -2,
		},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			name: "Zero add positive is positive",
			a:    0, b: 1, want: 1,
		},
		{
			name: "Positive add positive is positive",
			a:    2, b: 2, want: 4,
		},
		{
			name: "Negative add positive of same value is zero",
			a:    -10, b: 10, want: 0,
		},
		{
			name: "Zero add fraction is that fraction",
			a:    0, b: 5.5, want: 5.5,
		},
		{
			name: "Negative add positive of same fraction is zero",
			a:    -0.5, b: 0.5, want: 0,
		},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{
			name: "Positive minus a smaller positive is positive",
			a:    4, b: 2, want: 2,
		},
		{
			name: "Positive minus a bigger positive is negative",
			a:    2, b: 4, want: -2,
		},
		{
			name: "Positive minus a negative is positive",
			a:    2, b: -4, want: 6,
		},
		{
			name: "Negative minus a smaller negative is negative",
			a:    -4, b: -2, want: -2,
		},
		{
			name: "Positive fraction minus a smaller positive is a positive fraction",
			a:    4.5, b: 2, want: 2.5,
		},
		{
			name: "Negative minus positive is negative",
			a:    -4.5, b: 2, want: -6.5,
		},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Subtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
