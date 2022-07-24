package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Nur"
	names[1] = "Awwabin"
	names[2] = "Yoga Saputra"

	fmt.Println(names)
	fmt.Println("<=====>")

	var values = [3]int{
		90,
		86,
		98,
	}

	fmt.Println(values)

	fmt.Println("<=====>")
	// len(array) => untuk mendapatkan panjang array
	// array[index] => mendapatkan data diposisi index
	// array[index] = value => mengubah data di posisi index

	fmt.Println(len(names))
	fmt.Println(len(values))

}
