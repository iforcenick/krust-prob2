package main

import (
	"fmt"
	"testing"
)

type Pair struct {
	input  int
	output int
}

func TestGetSum(t *testing.T) {
	arr := []Pair{
		{10, 23},
		{1000, 233168},
	}
	fmt.Print(arr)
	for _, pair := range arr {
		test_name := fmt.Sprintf("%d primetest : ", pair.input)
		t.Run(test_name, func(t *testing.T) {
			ans := getSum(pair.input)
			if ans != pair.output {
				t.Errorf("got %d, want %d", ans, pair.output)
			}
		})
	}
}
