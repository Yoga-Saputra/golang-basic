package main

import (
	"fmt"
	"time"
)

/*
	Package time

	Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
	https://golang.org/pkg/time/

	Beberapa Function di Package time :

	- time.Now() => Untuk mendapatkan waktu saat ini
	- time.Date(...) => Untuk membuat waktu
	- time.Parse(layout, string) => Untuk memparsing waktu dari string


*/

func main() {
	now := time.Now()

	fmt.Println(now.Local())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	fmt.Println("<==== UTC ====>")

	UTC := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(UTC.Local())

	fmt.Println("<==== Parse ====>")
	parse, _ := time.Parse(time.RFC3339, "2022-01-02T15:04:05Z")
	fmt.Println(parse)
}
