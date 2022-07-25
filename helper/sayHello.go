package helper

import "fmt"

var version = "1.0.0" //tidak bisa diakses dari luar
var Application = "Go-Lang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("GoodBye ", name)
}
