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

type Customer struct {
	Nama   string
	HP     string
	Alamat []string
}

// function -> bebas diakses selama parameternya sesuai
//
//	func TambahQty(b *Barang, qtyTambahan int) {
//		b.Qty += qtyTambahan
//	}
//
// Method -> fungsi khusus yang dimiliki oleh sebuah struct
func (b *Barang) TambahQty(qtyTambahan int) {
	b.Qty += qtyTambahan
}

//Varian
//warna
//expired
//gambar -> teknis, biar penjelasan ga panjang

func main() {
	barang1 := Barang{1, "Royco", 1000, "Penyedap makanan bikin pinter", 100, 1}
	barang2 := Barang{Harga: 500, ID: 2, Nama: "Masako", Qty: 100, Berat: 1}
	var barang3 = new(Barang)
	barang3.Nama = "Garam"
	var barang4 Barang
	barang4.Nama = "kopi"
	var barang5 = Barang{}
	barang5.Nama = "Lada"
	fmt.Println(barang1, barang2, barang3, barang4)
	cus1 := Customer{"Indra", "081234", []string{"Lampung", "jakarta", "Medan"}}
	cus2 := Customer{}
	cus2.Nama = "Jerry"
	cus2.HP = "1234567"
	cus2.Alamat = []string{
		"pasuruan",
		"surabaya",
		"malang",
	}

	listCustomer := []Customer{}
	listCustomer = append(listCustomer, cus1)
	listCustomer = append(listCustomer, cus2)

	fmt.Println(cus1)
	fmt.Println(cus1.Alamat[1])
	fmt.Println(cus2)
	fmt.Println("Print list customer", listCustomer)
	fmt.Println("Print dari sliceCustomer", listCustomer[0].Nama, listCustomer[0].Alamat[0])
	// PrintPerkenalan(cus2)

	// Local struct
	// type tmp struct {
	// 	nama string
	// 	hp   string
	// }

	// t := tmp{}
	// t.hp = "0812345"
	fmt.Println("==========")
	fmt.Println("Qty barang 1 sebelum nambah", barang1.Qty)
	barang1.TambahQty(10)
	fmt.Println("Qty barang 1 setelah nambah", barang1.Qty)

	// function method
	fmt.Println("================")
	PrintPerkenalan("putra")
	PrintPerkenalan("fahmi")
	PrintPerkenalan(cus2.Nama)
	PrintPerkenalan(barang1.Nama)

	// yang kita pengen, perkenalan HANYA UNTUK CUSTOMER
	cus2.Kenalan()
	cus1.Kenalan()
}

// function
func PrintPerkenalan(nama string) {
	fmt.Println("Halo saya", nama)
}

// method -> fungsi khusus menempel ke sebuah struct
func (c Customer) Kenalan() {
	fmt.Println("Halo saya", c.Nama, "HP saya", c.HP)
}
