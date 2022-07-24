package main

import "fmt"

type Animal struct {
	Name string
}

type HasName interface {
	GetName() string
}

// struct method
func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello Push ", hasName.GetName())
}

func main() {
	animal := Animal{Name: "caty"}
	SayHello(animal)

}
