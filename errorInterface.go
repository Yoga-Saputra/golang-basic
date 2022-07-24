package main

import (
	"errors"
	"fmt"
)

// Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error , nama interface ya adalah Error

func Pembagi(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		// errors bawaan dari golang
		return 0, errors.New("Pembagian dengan nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagi(100, 0)

	if err == nil {
		fmt.Println("Hasil ", hasil)
	} else {
		fmt.Println("Error ", err.Error())
	}
}
