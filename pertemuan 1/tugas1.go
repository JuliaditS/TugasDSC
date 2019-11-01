package main

import "fmt"

func main() {
	const Nim string = "10119147"
	const Pass string = "01072001"
	var usernim, userpass string

	Salah := 0

	for Salah < 3 {
		fmt.Print("Masukan NIM : ")
		fmt.Scanf("%s", &usernim)
		fmt.Print("Masukan Password : ")
		fmt.Scanf("%s", &userpass)

		if (usernim == Nim) && (userpass == Pass) {
			fmt.Println("Benar")
			Salah = 4
		} else {
			fmt.Println("Salah")
			Salah++
		}

	}
	if Salah == 3 {
		fmt.Println("Blokir")
	}
}
