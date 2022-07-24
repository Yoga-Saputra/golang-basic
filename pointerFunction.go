package main

import "fmt"

/*
	Pointer di function
	- Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
	- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
	- Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
	- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
	- Untuk melakukan ini, kita juga bisa menggunakan pointer di function
	- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya
*/

type Address struct {
	City, Province, Country string
}

func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{
		City:     "Blora",
		Province: "Jawa Tengah",
		Country:  "",
	}
	changeAddressToIndonesia(&address)

	fmt.Println(address) //tidak berubah
}
