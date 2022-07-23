package main

import "fmt"

func main() {
	var name string

	name = "Nur Awwabin Yoga Saputra"

	fmt.Println(name)

	var callme = "Yoga"

	fmt.Println(callme)

	var age int8 = 30
	fmt.Println(age)

	// var bisa diganti dengan :=
	country := "indonesia"
	fmt.Println(country)

	country = "Singapore"
	fmt.Println(country)

	var (
		firstName = "Yoga"
		lastName  = "Saputra 123"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
