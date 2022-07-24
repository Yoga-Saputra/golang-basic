package main

import "fmt"

func main() {
	// recover adalah function yang bisa kita gunakan untuk menangkap data panic
	// dengan recover proses panic akan berhenti, sehingga program akan tetap berjalan
	runApp(true)
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message ", message)
	}
	fmt.Println("APlikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR Panic")
	}
}
