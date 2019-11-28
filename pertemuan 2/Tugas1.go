package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pengguna struct {
	nim      string
	password string
	nama     string
}

var scanner = bufio.NewReader(os.Stdin)
var listPengguna = []Pengguna{}
var menu int

func main() {
	for menu != 7 {
		inputMenu()
		switch menu {
		case 1:
			TampilData()
		case 2:
			tambahData()
		case 3:
			deleteData()
		case 4:
			updateData()
		case 5:
			cariData()
		case 6:
			login()
		case 7:
			fmt.Println("Selamat jalan ~")
		default:
			fmt.Println("Menu yang dimasukkan tidak tersedia ~")
		}
	}
}

func inputMenu() {
	fmt.Println("\n1. Lihat data")
	fmt.Println("2. Tambah data")
	fmt.Println("3. Hapus data")
	fmt.Println("4. Ubah data")
	fmt.Println("5. Cari data")
	fmt.Println("6. Login  ")
	fmt.Println("7. Keluar")
	fmt.Print("Masukkan menu yang dipilih: ")
	fmt.Scanln(&menu)
	fmt.Println("\n")
}

func TampilData() {
	for _, pengguna := range listPengguna {
		fmt.Print("NIM : ", pengguna.nim)
		fmt.Print("NAMA : ", pengguna.nama)
		fmt.Println("PASSWORD : ", pengguna.password)
	}
}

func tambahData() {
	fmt.Print("Masukkan Nim: ")
	nim, _ := scanner.ReadString('\n')
	fmt.Print("Masukkan Password: ")
	password, _ := scanner.ReadString('\n')
	fmt.Print("Masukkan Nama: ")
	nama, _ := scanner.ReadString('\n')

	penggunaBaru := Pengguna{nim, password, nama}
	listPengguna = append(listPengguna, penggunaBaru)
}

func deleteData() {
	fmt.Print("Masukkan Nim: ")
	nim, _ := scanner.ReadString('\n')

	ketemu := false
	for index, pengguna := range listPengguna {
		if pengguna.nim == nim {
			listPengguna = append(listPengguna[:index], listPengguna[index+1:]...)
			ketemu = true
			break
		}
	}

	if ketemu {
		TampilData()
	} else {
		fmt.Print("Tidak ketemu")
	}
}

func updateData() {
	fmt.Print("Masukkan Nim: ")
	nim, _ := scanner.ReadString('\n')

	ketemu := false
	for index, pengguna := range listPengguna {
		if pengguna.nim == nim {

			fmt.Print("Masukkan Nama: ")
			nama, _ := scanner.ReadString('\n')
			fmt.Print("Masukkan Password: ")
			password, _ := scanner.ReadString('\n')

			if nama != "" {
				listPengguna[index].nama = nama
			}

			if password != "" {
				listPengguna[index].password = password
			}

			ketemu = true
			break
		}
	}

	if ketemu {
		TampilData()
	} else {
		fmt.Print("Tidak ketemu")
	}
}

func cariData() {
	fmt.Print("Masukkan Nim: ")
	nim, _ := scanner.ReadString('\n')

	ketemu := false
	penggunaYangDicari := Pengguna{}
	for _, pengguna := range listPengguna {
		if pengguna.nim == nim {
			penggunaYangDicari = pengguna
			ketemu = true
			break
		}
	}

	if ketemu {
		fmt.Print("Nim: ", penggunaYangDicari.nim)
		fmt.Print("Nama: ", penggunaYangDicari.nama)
		fmt.Print("Password: ", penggunaYangDicari.password)
	} else {
		fmt.Print("Tidak ketemu")
	}
}

func login() {
	fmt.Print("Masukkan Nim: ")
	nim, _ := scanner.ReadString('\n')

	ketemu := false
	for _, pengguna := range listPengguna {
		if pengguna.nim == nim {
			ketemu = true
			fmt.Print("Masukkan Password: ")
			password, _ := scanner.ReadString('\n')
			if pengguna.password == password {
				fmt.Println("Sukses login")
			} else {
				fmt.Println("Gagal login")
			}

			return
		}
	}

	if !ketemu {
		fmt.Println("Pengguna tidak ditemukan")
	}
}