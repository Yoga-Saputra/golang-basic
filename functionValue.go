package main

import "fmt"

func main() {
	goodBye := getGoodBye("Yoga")
	fmt.Println(goodBye)
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
