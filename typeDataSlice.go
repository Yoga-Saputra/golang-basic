package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"Februari",
		"Maret",
		"April",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"November",
		"Desember",
	}

	// membuat slice dari array
	// array[low:high] => membuat slice dari array dimulai index low sampai index sebelum high
	// array[low:] => membuat slice dari array dimulai index low sampai index akhir  di array
	// array[:high] => membuat slice dari array dimulai index 0 sampai index sebelum high
	// array[:] => membuat slice dari array dimulai index 0 sampai index akhir di array

	// function slice
	// len(slice) => untuk mendapatkan panjang
	// cap(slice) => untuk mendapatkan kapasitas
	// append(slice, data) => membuat slice baru dengan menambahkan data keposisi akhir terakhir slice, jika kapasitas sudah penuh, makan akan membuat array baru
	// make([]TypeData, lenght, capacity) => membuat slice baru
	// copy(destination, source) => menyalin slice dari source ke destination

	// mengambil array ke 4 - 7
	slice1 := months[4:7]
	fmt.Println(months)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	// fmt.Println(append(slice1))

	fmt.Println("<=====>")

	// merubah isi array index ke 5
	months[5] = "Diubah"
	fmt.Println(months)

	// merubah isi slice1 pada index ke 0 dan di array pun ikut ke ibah
	slice1[0] = "UbahSlice"
	fmt.Println(months)

	fmt.Println("<==END ==>")

	// append slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "SeninBaru"
	daySlice1[1] = "SelasaBaru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	fmt.Println("<==APPEND==>")

	daySlice2 := append(daySlice1, "LiburBaru")
	daySlice2[0] = "Ops"
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("<==Make Slice==>")
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Yoga"
	newSlice[1] = "Yoga"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("<==Copy Slice==>")

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

}
