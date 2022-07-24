package main

import "fmt"

func main() {
	// recursive function adalah function yang memanggil function dirinya sendiri

	fmt.Println("<==== factorial for loop ====>")
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)

	fmt.Println("<==== factorial recursive function ====>")
	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}

// factorial for loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

// recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
