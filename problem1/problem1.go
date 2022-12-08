package main

import "fmt"

// anggap struct adalah tipe data custom, yang dapat mengakomodir keperluan kalian
type Kelas struct {
	// property
	murid []string
	nilai []int
}

// kalau struct memerlukan fungsi khusus silahkan buat sebuah method untuk struct tsb
func (k Kelas) Avarage() float64 {
	sum := 0
	var avg float64
	for i := 0; i < len(k.nilai); i++ {
		sum += k.nilai[i]

	}
	avg = float64(sum) / float64(len(k.nilai))
	return avg
}

func (k Kelas) Min() (min int, name string) {
	min = k.nilai[0]
	for i, v := range k.nilai { // elemen pertama?
		if v < min {
			min = v
			name = k.murid[i]
		}
	}
	return min, name
}

func (k Kelas) Max() (max int, name string) {
	for i, v := range k.nilai {
		if max < v {
			max = v
			name = k.murid[i]
		}
	}
	return max, name
}

func main() {
	var a = Kelas{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input ", i, " Student’s Name : ")
		fmt.Scan(&name)
		a.murid = append(a.murid, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.nilai = append(a.nilai, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : ", nameMax, " (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : ", nameMin, " (", scoreMin, ")")

	// var b = Kelas{}
	// b.Avarage() -> tidak ada yang diproses

	// APAKAH a.murid == b.murid (?) TIDAK SAMA
}
