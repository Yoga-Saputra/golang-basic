package main

import "fmt"

func main() {
	// panic function adalah function yang bisa kita gunakan untuk menghentikan program
	// panic function biasa dipanggil ketika terjadi error pada saat program kita berjalan
	// saat panic function dipanggil, program akan berhenti, tetapi defer function tetap dieksekusi
	runApp(true)
}

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}

	fmt.Println("Aplikasi berjalan")
}
