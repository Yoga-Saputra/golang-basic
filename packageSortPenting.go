package main

import (
	"fmt"
	"sort"
)

/*
	Package sort

	Package sort adalah package yang berisikan utilitas untuk proses pengurutan
	Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
	https://golang.org/pkg/sort/

	sort.Interface
	type Interface Interface {
		// Len is the number one of elements in the collection
		Len() int

		// Less reports whether the element with
		// index i should sort before he element with index j
		Less(i, j int) bool

		// Swap swaps the element with index i and j
		//Swap(i, j int)
	}
*/

type User struct {
	Name string
	Age  int
}

// type alias of array user
type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

// sort name
func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Name < userSlice[j].Name
}

// sort age
// func (userSlice UserSlice) Less(i, j int) bool {
// 	return userSlice[i].Name < userSlice[j].Name
// }

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"Yoga", 24},
		{"Ajeng", 26},
		{"Kay", 5},
	}
	fmt.Println("Sebelum di sort")
	fmt.Println(users)

	sort.Sort(UserSlice(users))

	fmt.Println("Sesudah di sort")
	fmt.Println(users)
}
