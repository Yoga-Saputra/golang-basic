package main

import (
	"container/list"
	"fmt"
)

/*
	Package container/list

	Package container/list adalah implementasi struktur data double linked list di Go-Lang
	https://golang.org/pkg/container/list/

*/

func main() {
	data := list.New()
	data.PushBack("Yoga")
	data.PushBack("Saputra")
	data.PushBack("Kay")

	// dari depan ke belakang
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
