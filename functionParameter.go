package main

import "fmt"

func main() {
	sayHello("Yoga", "Saputra")
}

func sayHello(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}
