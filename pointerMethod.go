package main

import "fmt"

/*
	Pointer di method
	- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
	- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
*/

type Man struct {
	Name string
}

// struct method
// pointer
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	yoga := Man{"Yoga"}
	yoga.Married()

	fmt.Println(yoga.Name)
}
