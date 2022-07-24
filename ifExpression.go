package main

import "fmt"

func main() {
	name := "Yoga"

	if name == "yoga" {
		fmt.Println("Hello Yoga")
	} else if name == "Yoga" {
		fmt.Println("Ini benar Yoga")
	} else {
		fmt.Println("No yoga")
	}

	fmt.Println("<==== if short statement ====>")
	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
