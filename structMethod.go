package main

import "fmt"

type CustomerPremium struct {
	Name, Title string
	Age         int
}

/*
	struct adalah tipe data seperti tipe data lainnya , dia bisa digunakan sebagai parameter untuk function.
	Namum jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function.
	method adalah function
*/
func main() {
	yoga := CustomerPremium{Title: "Backend Developer"}
	yoga.sayHello()
}

func (customerPremium CustomerPremium) sayHello() {
	fmt.Println("Hello, I'm ", customerPremium.Title)
}
