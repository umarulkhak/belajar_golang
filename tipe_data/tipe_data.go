package main

import "fmt"

func main() {
	// tipe data numerik non-decimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// tipe data numerik non-decimal
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//Tipe Data bool (Boolean)
	var exist bool = true
	fmt.Printf("exist? %t \n,", exist)

	//Tipe Data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var message1 = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message1)

}
