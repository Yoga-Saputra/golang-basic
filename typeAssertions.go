package main

import "fmt"

/*
	Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
	fitur ini sering digunakan ketika kita bertemu dengan data interface kosong
*/

func random() interface{} {
	return "tes"
}

func main() {
	result := random()

	fmt.Println("<====switch expression for type assertions====>")

	/*
		Type Assertions menggunakan switch
		agar lebih aman dalam menggunakan type assertions, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	*/
	switch value := result.(type) {
	case string:
		fmt.Println("String ", value)
	case int:
		fmt.Println("Int ", value)
	case bool:
		fmt.Println("Bool ", value)
	default:
		fmt.Println("Unkown")

	}

	// resultString := result.(string)
	// fmt.Println(resultString)

	// fmt.Println("<====Panic====>")
	// resultInt := result.(int) //panic
	// fmt.Println(resultInt)

}
