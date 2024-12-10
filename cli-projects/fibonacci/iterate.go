package fibonacci

// Function to generate Fibonacci sequence up to n terms
func Iterate(n int) []int {
	// Slice to store the Fibonacci sequence
	fibonacci := make([]int, n)

	// First two numbers in the Fibonacci sequence
	if n > 0 {
		fibonacci[0] = 0
	}
	if n > 1 {
		fibonacci[1] = 1
	}

	// Generate the Fibonacci sequence
	for i := 2; i < n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	return fibonacci
}
