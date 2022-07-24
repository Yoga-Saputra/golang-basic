package main

import "fmt"

func main() {
	total := sumAll(10, 10, 10, 10, 5)
	fmt.Println(total)

	slice := []int{10, 20, 30}
	total = sumAll(slice...)
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
