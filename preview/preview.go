package main

import "fmt"

type Barang struct {
	ID        int // TERSERAH!! SESUAI KEBUTUHAN SAJA
	Nama      string
	Harga     int
	Deskripsi string
	Qty       int
	Berat     int
}

func main() {
	var listBarang []Barang
	var input int
	for input != 9 {
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Lihat Barang")
		fmt.Println("9. EXIT")
		fmt.Scanln(&input)

		if input == 1 {
			barangBaru := Barang{}
			fmt.Scanln(&barangBaru.ID)
			fmt.Scanln(&barangBaru.Nama)
			fmt.Scanln(&barangBaru.Harga)
			fmt.Scanln(&barangBaru.Berat)
			listBarang = append(listBarang, barangBaru)
		} else if input == 2 {
			fmt.Println(listBarang)
		}
	}
}
