package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// length := len(numbersToSum)
	// ans := make([]int, length)
	// for idx, numbers := range numbersToSum {
	// 	ans[idx] = Sum(numbers)
	// }

	var ans []int
	for _, numbers := range numbersToSum {
		ans = append(ans, Sum(numbers))
	}
	return ans
}

func SumAllTails(numberTails ...[]int) []int {
	var sums []int
	for _, numbers := range numberTails {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}
