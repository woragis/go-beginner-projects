package fibonacci

import "fmt"

func Init() {
	var n int

	// Ask the user for the number of terms
	fmt.Print("Enter the number of Fibonacci terms to generate: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Please enter a valid positive number.")
		return
	}

	// Generate Fibonacci sequence
	fibonacci := Iterate(n)

	// Display the Fibonacci sequence
	fmt.Println("Fibonacci sequence:")
	for _, num := range fibonacci {
		fmt.Printf("%d ", num)
	}

	// Display the Fibonacci sequence using recursion
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", Recursive(i))
	}
	fmt.Println()
	
}
