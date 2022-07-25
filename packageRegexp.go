package main

import (
	"fmt"
	"regexp"
)

/*
	Package regexp

	-	Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
	- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
	-	https://github.com/google/re2/wiki/Syntax
	-	https://golang.org/pkg/regexp/

	Beberapa Function di Package regexp :
	-	regexp.MustCompile(string) => Membuat Regexp
	-	Regexp.MatchString(string) bool => Mengecek apakah Regexp match dengan string
	- Regexp.FindAllString(string, max) => Mencari string yang match dengan maximum jumlah hasil
*/

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eyo"))

	search := regex.FindAllString("eko, eto, eco, eyo", 2)
	fmt.Println(search)
}
