package main

import (
	"fmt"
	"strings"
)

/*
	Package String
	- strings.Trim(string, cutset) => Memotong cutset di awal dan akhir string
	- strings.ToLower(string) => Membuat semua karakter string menjadi lower case
	- strings.ToUpper(string) => Membuat semua karakter string menjadi upper case
	- strings.Split(string, separator) => Memotong string berdasarkan separator
	- strings.Contains(string, search) => Mengecek apakah string mengandung string lain
	- strings.ReplaceAll(string, from, to) => Mengubah semua string dari from ke to

*/

func main() {
	fmt.Println(strings.Contains("Yoga Saputra", "Yoga"))
	fmt.Println(strings.Split("Yoga Saputra", ""))
	fmt.Println(strings.ToLower("Yoga Saputra"))
	fmt.Println(strings.ToUpper("Yoga Saputra"))
	fmt.Println(strings.Trim(" Yoga Saputra ", " "))
	fmt.Println(strings.ReplaceAll("Yoga Saputra", "yoga", "Kay"))
}
