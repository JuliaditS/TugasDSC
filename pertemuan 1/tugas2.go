package main

import "fmt"

func main() {
	var a, b float32
	Menu := 0

	for Menu != 5 {
		fmt.Println("Kalkulator sederhana")
		fmt.Println("1.Tambah")
		fmt.Println("2.Kurang")
		fmt.Println("3.Kali")
		fmt.Println("4.Bagi")
		fmt.Println("5.Keluar")
		fmt.Print("Masukan Menu : ")
		fmt.Scan(&Menu)

		for Menu < 1 || Menu > 5 {
			fmt.Println("\nMasukkan menu dengan benar!!")
			fmt.Print("Masukan Menu : ")
			fmt.Scan(&Menu)
		}

		if Menu >= 1 && Menu < 5 {
			fmt.Print("Masukan Angka pertama : ")
			fmt.Scan(&a)
			fmt.Print("Masukan Angka kedua : ")
			fmt.Scan(&b)
		}

		switch Menu {
		case 1:
			fmt.Println("Hasilnya adalah : ", tambah(a, b), "\n")
		case 2:
			fmt.Println("Hasilnya adalah : ", kurang(a, b), "\n")
		case 3:
			fmt.Println("Hasilnya adalah : ", kali(a, b), "\n")
		case 4:
			fmt.Println("Hasilnya adalah : ", bagi(a, b), "\n")
		}
	}
}

func tambah(a, b float32) float32 {
	return a + b
}

func kurang(a, b float32) float32 {
	return a - b
}

func kali(a, b float32) float32 {
	return a * b
}

func bagi(a, b float32) float32 {
	return a / b
}
