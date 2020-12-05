package main

// Sum sum
func Sum(numbers []int) int {
	sum := 0
	// ★for文
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	// ★range
	// rangeも配列の反復処理を行う。
	// index,中身が返却される
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll SumAll
func SumAll(numbersToSum ...[]int) []int {

	// lengthOfNumbers := len(numbersToSum)

	// make sliceの作成
	// sums := make([]int, lengthOfNumbers)

	var sums []int
	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)

		// ★append関数
		// sliceと新しい値を受け取り、その中にあるすべての項目を含む、新しいsliceを返す
		sums = append(sums, Sum(numbers))
	}

	return sums
}
