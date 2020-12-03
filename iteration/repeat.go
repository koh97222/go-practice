package iteration

// Repeat 繰り返し
func Repeat(character string, cnt int) string {
	var repeated string

	for i := 0; i < cnt; i++ {
		// repeated = repeated + character
		repeated += character
	}
	return repeated
}
