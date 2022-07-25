package main

import (
	"fmt"
	"reflect"
)

/*
	Package reflect

	-	Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
	- Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
	-	Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa eksplorasi package reflec ini secara otodidak
	-	Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
	-	https://golang.org/pkg/reflect/
*/

type Sample struct {
	Name string
}

type StructTagSample struct {
	Name string `required:"true" max:"10"`
}

func main() {

	// Kode Program Package reflect
	sample := Sample{"Yoga"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println(sampleType.NumField())
	fmt.Println(structField.Name)

	fmt.Println("<=== Kode Program StructTag ===>")
	sampleTag := StructTagSample{"Yoga"}
	sampleTypeTag := reflect.TypeOf(sampleTag)
	structFieldTag := sampleTypeTag.Field(0)
	name := structFieldTag.Name
	required := structFieldTag.Tag.Get("required")
	max := structFieldTag.Tag.Get("max")
	fmt.Println(name)
	fmt.Println(required)
	fmt.Println(max)

}
