package main

import (
	"fmt"
	"math"
)

func getSum(n int) int {
	n -= 1

	cnt3 := math.Floor(float64(n) / 3)
	cnt5 := math.Floor(float64(n) / 5)
	cnt15 := math.Floor(float64(n) / 15)
	sum3 := 3 * (cnt3 * (cnt3 + 1) / 2)
	sum5 := 5 * (cnt5 * (cnt5 + 1) / 2)
	sum15 := 15 * (cnt15 * (cnt15 + 1) / 2)

	return int(sum3 + sum5 - sum15)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(getSum(n))
}
