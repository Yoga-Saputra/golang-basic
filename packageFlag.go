package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your host")

	flag.Parse()

	fmt.Println(*host)
}
