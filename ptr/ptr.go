package main

import "fmt"

func main() {
	var a string = "jery"
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b = "thomas"
	fmt.Println(a)
	fmt.Println(b)

	var c *string // variable ini menerima nilai berupa alamat memori dari data dengan tipe sejnis
	c = &a
	fmt.Println(c)  // isi dari c adalah alamat dari a
	fmt.Println(&c) // alamatnya c
	fmt.Println("Mengakses nilai dari pointer", *c)
	*c = "Tono"
	fmt.Println("INI A!!", a)
	fmt.Println(b)
	fmt.Println("Setelah c diubah", *c)
	fmt.Println("==============================")
	a = "Budi"
	fmt.Println("INI A!!", a)
	fmt.Println(b)
	fmt.Println("Setelah c diubah", *c)
	fmt.Println("==============================")
	var d *string
	d = c
	fmt.Println("Ini isi dari d awal ", *d)
	*d = "kasino"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*c)
	fmt.Println(*d)

	// d := new(int)
}
