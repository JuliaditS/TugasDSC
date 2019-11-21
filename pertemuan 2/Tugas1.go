package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	nim      string
	password string
	nama     string
}

type MenuPilihan interface {
	TampilData()
}

var scanner = bufio.NewReader(os.Stdin)
var listUser = []User{}

func main() {
	var menu MenuPilihan
	var NIM, Password, Nama string
	Menu := 0

	for Menu != 7 {
		fmt.Println("\n\nMenu Pilihan")
		fmt.Println("")
		fmt.Println("1. Tambah data")
		fmt.Println("2. lihat data")
		fmt.Println("3. Hapus data")
		fmt.Println("4. Ubah data")
		fmt.Println("5. Cari data")
		fmt.Println("6. Login")
		fmt.Println("7. Keluar")
		fmt.Print("Masukkan Menu : ")
		fmt.Scan(&Menu)
		fmt.Println("\n")

		for Menu < 1 || Menu > 7 {
			fmt.Println("\nMasukkan menu dengan benar!!")
			fmt.Print("Masukan Menu : ")
			fmt.Scan(&Menu)
			fmt.Println("\n")
		}

		switch Menu {
		case 1:
			fmt.Print("Masukkan Nim: ")
			NIM, _ = scanner.ReadString('\n')

			fmt.Print("Masukkan Password: ")
			Password, _ = scanner.ReadString('\n')

			fmt.Print("Masukkan Nama: ")
			Nama, _ = scanner.ReadString('\n')

			newUser := User{NIM, Password, Nama}
			listUser = append(listUser, newUser)
			fmt.Println(listUser)
		case 2:
			menu = User{NIM, Password, Nama}
			menu.TampilData()
		case 3:
			HapusData()
		case 4:
			UbahData()
		case 5:
			CariData()
		case 6:
			Login()
		}
	}
}

func (User) TampilData() {
	listuser := User{}
	var NIM, Password, Nama string
	for _, user := range listUser {
		listuser = user
		NIM = listuser.nim
		Password = listuser.password
		Nama = listuser.nama
		fmt.Print("NIM : ", NIM)
		fmt.Print("NAMA : ", Nama)
		fmt.Println("PASSWORD : ", Password)
	}
}

func HapusData() {

}

func UbahData() {

}

func CariData() {

}

func Login() {

}
