package main

import "fmt"

/*
	data nil adalah data kosong di golan
	nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan chann el
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Yoga")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}
}
