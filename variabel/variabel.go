package main

import "fmt"

func main() {
	var firstName string = "Umar"

	var lastName string
	lastName = "Ulkhak"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// fungsi fmt.Printf() digunakan untuk menampilkan output
	//"halo %s %s!\n", karakter %s di situ akan diganti dengan data string yang berada di parameter ke-2, ke-3, dan seterusnya.

	//Deklarasi Variabel Tanpa Tipe Data
	var firstName1 string = "john"
	//:= hanya digunakan sekali di awal pada saat deklarasi. Untuk assignment nilai selanjutnya harus menggunakan tanda =, contoh:
	// lastName := "wick"
	// lastName = "ethan"
	// lastName = "bourne"
	lastName1 := "wick"

	fmt.Printf("halo %s %s!\n", firstName1, lastName1)

	// //Deklarasi Multi Variabel
	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// //Pengisian nilai juga bisa dilakukan bersamaan pada saat deklarasi. Caranya dengan menuliskan nilai masing-masing variabel berurutan sesuai variabelnya dengan pembatas koma (,).
	// var fourth1, fifth1, string1 = "empat", "lima", "enam"
	// //atau lebih ringkas
	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"
}
