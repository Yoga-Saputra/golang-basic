package main

import "fmt"

type CustomerPremium struct {
	Name, Address string
	Age           int
}

func main() {
	// struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
	// struct biasanya representasi data dalam program aplikasi yang kita buat
	// data di struct biasanya di simpan dalam field
	// sederhananya struct adalah kumpulan dari field

	// data struct

	fmt.Println("<====== data struct ======>")
	var yoga CustomerPremium
	yoga.Name = "Yoga"
	yoga.Address = "Indonesia"
	yoga.Age = 24

	fmt.Println(yoga)

	fmt.Println("<====== struct literals ======>")
	Kay := CustomerPremium{
		Name:    "Kay",
		Address: "Blora",
		Age:     6,
	}
	fmt.Println(Kay)
}
