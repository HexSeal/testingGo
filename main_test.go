package main

import (
	"testing"
	"fmt"
)

// Go test flags:
// -v shows verbose test logs, -cover shows test coverage, -coverprofile=coverage.out shows you which files lack coverage

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

// go test -run=Calculate -bench=. 
// To run both test and benchmark

// go test -run=Bench -bench=.
// To just run benchmark

func benchmarkCalculate(input int, b *testing.B) {
    for n := 0; n < b.N; n++ {
        Calculate(input)
    }
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }

func TestOther(t *testing.T) {
    fmt.Println("Testing something else")
    fmt.Println("This shouldn't run with -run=calc")
}