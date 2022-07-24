package main

import "fmt"

// struct
type Person struct {
	Name string
}

// interface
type HasName interface {
	GetName() string
}

// struct method
func (person Person) GetName() string {
	return person.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello prgammer ", hasName.GetName())
}

/*
	interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
	sebuah interface berisikan definisi" method
	biasanya interface digunakan sebagai kontrak


	Implementasi interface :
	Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
	sehingga kita tidak perlu mengimplementasikan interface secara manual
*/
func main() {
	person := Person{Name: "Yoga"}
	SayHello(person)
}
