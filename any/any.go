package main

import (
	"fmt"
	"strconv"
)

func hitungUtang(angka1 string, angka2 int) any {
	cnv, err := strconv.Atoi(angka1)
	fmt.Println(err)
	return cnv + angka2
}

func main() {
	var x interface{}
	var z any

	x = []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	x = 10
	fmt.Println(x)
	x = "tono bermain bola"
	fmt.Println(x)

	z = map[string]int{"halo": 1}
	fmt.Println(z)

	fmt.Println(hitungUtang("baso", 15000))
	z = 100000
	var uangku int = 10000
	res := uangku + z.(int)
	fmt.Println("Hasil penjumlahan interface", res)

	z = []int{1, 2, 3, 4, 5}
	fmt.Println("Print z sebagi slice of int")
	for i := 0; i < len(z.([]int)); i++ {
		fmt.Println(z.([]int)[i])
	}
}
