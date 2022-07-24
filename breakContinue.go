package main

import "fmt"

func main() {
	// break digunakan untuk menghentikan seluruh perulangan
	fmt.Println("<==== break ====>")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}

	fmt.Println("<==== continue ====>")

	// continue digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan keperulangan selanjutnya
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}

}
