package main

import (
	"fmt"
	"sort"
)

func median(numbers []float64) float64 {
	// copy
	myCopy := make([]float64, len(numbers))
	copy(myCopy, numbers)

	// sorting
	sort.Float64s(myCopy)

	numbersLen := len(myCopy)
	position := numbersLen / 2

	if numbersLen%2 > 0 {
		return myCopy[position]
	}

	left := myCopy[position-1]
	right := myCopy[position]

	return (left + right) / 2
}

func main() {

	numbers := make([]float64, 4, 4)
	numbers[0] = 11
	numbers[1] = 55
	numbers[2] = 20
	numbers[3] = 100

	medianResult := median(numbers)
	fmt.Println(medianResult)

	numbers = make([]float64, 3, 3)
	numbers[0] = 11
	numbers[1] = 55
	numbers[2] = 20

	medianResult = median(numbers)
	fmt.Println(medianResult)
}
