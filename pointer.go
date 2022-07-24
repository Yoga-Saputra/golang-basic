package main

import "fmt"

/*
	Pass By Value
	- Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
	- artinya ,jika kita mengirim sebuah variable ke dalam function , method atau variable lain sebenernya yang dikirim adalah duplikasi value nya

	Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
	sederhananya dengan mengunakan pointer kita bisa membuat pass by reference

	Operator &  berfungsi untuk membuat sebuah variable dengan nilai pointer  ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variable nya

	Operator *
	Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
	Semua variable yang mengacu ke semua data yang sama tidak akan berubah
	Jika ingin mengubah seluruh variable yang mengacu ke data tersebut , kita bisa menggunakan operator *

	function new
	sebelumnya untuk membuat pointer dengan menggunakan operator &
	Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
	namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal

*/

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Blora", "Jawa Tengah", "Indonesia"}
	// dengan pointer
	var address3 *Address = &address1
	// tanpa pointer
	var address2 Address = address1

	address2.City = "Jakarta "
	address3.City = "Jogja "
	*address3 = Address{"semarang", "Jawa Tengah", "Indonesia"}

	fmt.Println(address1) //address1 berubah
	fmt.Println("<=== dengan pointer ===>")
	fmt.Println(address3)
	fmt.Println("<=== tanpa pointer ===>")
	fmt.Println(address2)
	fmt.Println("<=== function new ===>")
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Singapore"

	fmt.Println(alamat1) //alamat 1 berubah
	fmt.Println(alamat2)
}
