package main

import "fmt"

func main() {
	firstName, lastName, age := getFullName()
	firstNameBaru, _, _ := getFullName()
	fmt.Println(firstNameBaru)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}

func getFullName() (string, string, int8) {
	return "Yoga", "Saputra", 24
}
