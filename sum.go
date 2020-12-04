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
