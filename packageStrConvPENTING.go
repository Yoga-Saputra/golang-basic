package main

import (
	"fmt"
	"strconv"
)

/*
	strconv

	Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
	Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
	Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
	https://golang.org/pkg/strconv/

	Beberapa Function di Package strconv :

	- strconv.parseBool(string) => Mengubah string ke bool
	- strconv.parseFloat(string) => Mengubah string ke float
	- strconv.parseInt(string) => Mengubah string ke int64
	- strconv.FormatBool(bool) => Mengubah bool ke string
	- strconv.FormatFloat(float, … ) => Mengubah float64 ke string
	- strconv.FormatInt(int, … ) => Mengubah int64 ke string


*/

func main() {

	// string ke boolean
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error ", err.Error())
	}

	// string ke int64
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	//int64 ke string
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	// string ke int
	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)
}
