package main

import "fmt"

func main() {
	// constant adalah variable yg nilainya tidak bisa di ubah lagi setelah pertama kali di beri nilai
	// saat pembuatan constant, kita wajib langsung menginisialisasi datanya

	const (
		firstName string = "Yoga"
		lastName         = "Saputra 123"
		value            = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

}
