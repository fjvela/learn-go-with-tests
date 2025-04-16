package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	size := len(numbersToSum)
	result := make([]int, size)
	for i, numbers := range numbersToSum {
		result[i] = Sum(numbers)
	}
	return result
}
