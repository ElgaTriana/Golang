package main

import "fmt"

func main() {
	fmt.Println("HALLO DUNIA")

	var firstName string = "Elga"

	var lastName string = "Triana"

	var positif uint16 = 256

	var tinggi uint8 = 10

	var panjang uint8 = 2

	var lebar uint8 = 2

	var negatif = -12039403

	const namanih = "APA YAA?"

	fmt.Println("RUMUS APA YA INI LUPA")

	var hasil = panjang * lebar * tinggi

	fmt.Println("Hasilnya adalah = ", hasil)

	fmt.Println("END RUMUS APA YA INI LUPA")

	fmt.Println("john wick")

	fmt.Println("john", "wick")

	fmt.Print("john wick\n")

	fmt.Println("Nilainya : ", positif)

	fmt.Printf("Nilai 2 nya %d :", negatif)

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	fmt.Printf("Hai", firstName, lastName+"Kobra\n")

	fmt.Println("Hai", firstName, lastName+"!")

	for a := 0; a <= 10; a++ {
		if a%2 == 0 {
			fmt.Println("ANGKA GENAPNYA ADALAH :", a)
		} else {
			fmt.Println("ANGKA GANJILNYA ADALAH :", a)
		}
	}

	for b := 0; b < 10; b++ {
		for k := b; k < 10; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
