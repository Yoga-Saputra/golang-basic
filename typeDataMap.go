package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Yoga",
		"address": "Blora",
	}

	// add data in map
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	fmt.Println("<===function map ===>")
	// len(map)   										: untuk mendapatkan jumlah data di map
	// map(key)   										: mengambil data di map dengan key
	// map(key) = value								: mengubah data di map dengan key
	// make(map[TypeKey],TypeValue)   : membuat map baru
	// delete(map, key) 							: menghapus data di map dengan key

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Yoga"
	book["wrong"] = "Ops"

	fmt.Println(len(book))

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))

}
