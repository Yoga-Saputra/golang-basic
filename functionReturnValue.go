package main

import "fmt"

func main() {
	data := getHello("Yoga")
	fmt.Println(data)
	fmt.Println(getHello(""))
}

func getHello(name string) string {
	if name == "" {
		return "Hello Bro !!"
	} else {
		return "Hello " + name
	}
}
