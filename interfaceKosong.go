package main

import "fmt"

// interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya

/*
	implementasi interface kosong di golang :
	fmt.Println(a ...interface{})
	panic(v interface{})
	recover()interface{}

*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "UPS"
	}
}

func main() {
	kosong := Ups(2)
	fmt.Println(kosong)
}
