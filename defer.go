package main

import "fmt"

func main() {
	// defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
	// note : defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
	runApplication(10)
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	fmt.Println("Run Application")

	defer logging()

	result := 10 / value
	fmt.Println(result)

}
