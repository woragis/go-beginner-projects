package evenoddchecker

import "fmt"

func Check(num int) string {
	if (num % 2 == 0) {
		fmt.Printf("The number '%d' is: even", num)
		return "even"
	}
	fmt.Printf("The number '%d' is: odd", num)
	return "odd"
}