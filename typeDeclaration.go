package main

import "fmt"

func main() {
	// type data declaration adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada

	type NoKtp string
	type Married bool

	var noKtpYoga NoKtp = "1231231"
	var MarriedStatus Married = true

	fmt.Println(noKtpYoga)
	fmt.Println(MarriedStatus)
}
