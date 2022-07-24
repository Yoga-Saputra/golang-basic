package main

import "fmt"

func main() {

	var (
		a        = 10
		b        = 20
		c        = 20
		negative = -100
		positive = +100
	)
	c += 200

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("<=======>")
	fmt.Println(negative)
	fmt.Println(positive)

}
