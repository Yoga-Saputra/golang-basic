package main

import "fmt"

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("eko", blacklist)
	registerUser("anjing", blacklist)
	registerUser("admin", func(name string) bool {
		return name == "admin"
	})
}

type BlackListUser func(string) bool

func registerUser(name string, blacklist BlackListUser) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }
