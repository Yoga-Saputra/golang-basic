package main

import (
	"fmt"
	"math"
)

/*
	Package math
	Package math merupakan package yang berisikan constant dan fungsi matematika
	https://golang.org/pkg/math/

	function di package math
	- math.Round(float64) => Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat
	- math.Floor(float64) => Membulatkan float64 kebawah
	- math.Ceil(float64) => Membulatkan float64 keatas
	- math.Max(float64, float64) => Mengembalikan nilai float64 paling besar
	- math.Min(float64, float64) => Mengembalikan nilai float64 paling kecil
*/

func main() {
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))
	fmt.Println(math.Floor(1.7))
	fmt.Println(math.Ceil(1.3))

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}
