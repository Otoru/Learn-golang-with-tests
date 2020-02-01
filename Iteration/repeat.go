package iteration

// Repeat function returns the character repeatedly
func Repeat(character string, size int) string {
	var repeated string

	for i := 0; i < size; i++ {
		repeated += character
	}

	return repeated
}
