package main

import (
	"cli_games/projects/calculator"
	"fmt"
	"os"
)

func main() {
	var num1, num2 float64
	var operator string
	fmt.Println("Simple Calculator")
	fmt.Println("Enter first number:")

	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Error reading the first number:", err)
		os.Exit(1)
	}

	fmt.Println("Enter operator:")
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Error reading operator:", err)
		os.Exit(1)
	}

	fmt.Println("Enter second number:")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Error reading the second number:", err)
		os.Exit(1)
	}

	var result float64
	switch operator {
	case "+":
		result = calculator.Add(num1, num2)
	case "-":
		result = calculator.Subtract(num1, num2)
	case "*":
		result = calculator.Multiply(num1, num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed")
			return
		}
		result = calculator.Divide(num1, num2)
	default:
		fmt.Println("Invalid operator. Please use one of (+, -, *, /).")
		return
	}

	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
