package textreverser

func ReverseText(s string) string {
	// Convert string to a slice of runes to handle multi-byte characters
	runes := []rune(s)
	
	// Reverse the runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	// Convert the reversed rune slice back to a string
	return string(runes)
}

