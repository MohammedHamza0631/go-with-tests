package main

func Sum (numbers []int) int{
	sum := 0;
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)

	ans := make([]int, length)
	for idx, numbers := range numbersToSum {
		ans[idx] = Sum(numbers)
	}
	return ans
}