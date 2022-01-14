package main

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func Sum(numbers [5]int) int {
	// sum := 0
	// for _, number := range numbers {
		// sum += number
	// }
	// return sum
// }
