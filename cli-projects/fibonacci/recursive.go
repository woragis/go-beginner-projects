package fibonacci

// Recursive Fibonacci function
func Recursive(n int) int {
	if n <= 1 {
		return n
	}
	return Recursive(n-1) + Recursive(n-2)
}
