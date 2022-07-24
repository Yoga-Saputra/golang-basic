package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	fmt.Println("<==== for dengan statement ====>")
	// init statement => statement sebelum for dieksekusi
	// post statement => statement yg akan selalu di eksekusi diakhir tiap perulangan
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	// slice
	slices := []string{"Yoga", "Kay", "Aqmar"}
	for a := 0; a < len(slices); a++ {
		fmt.Println(slices[a])
	}

	fmt.Println("<==== for range ====>")

	for index, slice := range slices {
		fmt.Println("index ", index, " = ", slice)
	}
	// atau bisa seperti ini
	for _, slice := range slices {
		fmt.Println("slice  = ", slice)
	}

	// map
	persons := make(map[string]string)
	persons["name"] = "Yoga"
	persons["title"] = "Developer"

	for key, person := range persons {
		fmt.Println(key, " = ", person)
	}

}
