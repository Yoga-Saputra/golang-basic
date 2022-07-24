package main

import "fmt"

func main() {
	sayHelloWithFilter("Yoga", spamFilter)
	sayHelloWithFilter("Anjing ", spamFilter)
}

// type declaration
type filterType func(string) string

// filter
// func sayHelloWithFilter(name string, filter func(string) string) {
func sayHelloWithFilter(name string, filter filterType) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
