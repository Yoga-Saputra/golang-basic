package main

import "fmt"

func main() {
	name := "Yoga"

	switch name {
	case "yoga":
		fmt.Println("Hello yoga")
	case "Kay":
		fmt.Println("Ini adalah Kay")
	case "Yoga":
		fmt.Println("Ini benar Yoga")
	default:
		fmt.Println("Kamu siapa?")
	}

	fmt.Println("<==== switch short statement =====>")
	length := len(name)
	switch length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	fmt.Println("<==== switch tanpa kondisi =====>")

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
