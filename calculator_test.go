package calculator_test

import (
	"calculator"
	"fmt"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
	errExpected bool
	}


func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase {
		{ a: 2, b: 2, want: 4 },
		{ a: 1, b: 1, want: 2 },
		{ a: 5, b: 0, want: 5 },
		{ a: 0, b: 0, want: 0 },
	}
	for _, tc:= range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f ", tc.a, tc.b, tc.want, got)
		}

	} 

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase {
		{ a: 2, b: 1, want: 1 },
		{ a: 1, b: 1, want: 0 },
		{ a: 0, b: 1, want: -1 },
	}
	for _, tc:= range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}

	} 
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase {
		{ a: 2, b: 2, want: 4 },
		{ a: 1, b: 1, want: 1 },
		{ a: 5, b: 0, want: 0 },
		{ a: 2.5, b: 2.5, want: 6.25},
	}
	for _, tc:= range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}

	} 
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase {
		{ a: 2, b: 2, want: 1, errExpected: false },
		{ a: 1, b: 1, want: 1, errExpected: false },
		{ a: 2.5, b: 2.5, want:1, errExpected: false},
		{ a: 5, b: 0, want:0, errExpected: true},
	}
	for _, tc:= range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		fmt.Printf(" error: %s,", err)
		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
		if tc.errExpected !=  (err != nil) {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, err)
		}
	}
} 

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []testCase {
		{ a: 4, want: 2, errExpected: false },
		{ a: 1, want: 1, errExpected: false },
		{ a: -5, want: 0, errExpected: true},
	}
	for _, tc:= range testCases {
		got, err := calculator.Sqrt(tc.a)
		fmt.Printf(" error: %s,", err)
		if !tc.errExpected && tc.want != got {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		}
		if tc.errExpected !=  (err != nil) {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, err)
		}
	} 
}
