package main

import "fmt"

func main() {
	firstName, lastName, age := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}

func getCompleteName() (firstName string, lastName string, age int8) {
	firstName = "Yoga"
	lastName = "Saputra"
	age = 24

	return
}
