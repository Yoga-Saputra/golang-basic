package main

import "fmt"

func main() {
	// closure adalah kemampuan sebuah function berinteraksi dengan data data disekitarnya dalam scope yang sama
	// harag gunakan fitur closur ini dengan bijak saat kita membuat aplikasi
	name := "Yoga"
	counter := 0
	increment := func() {
		name = "Kay"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
