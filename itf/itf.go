package main

type IT struct{}

func (i *IT) Coding()       {}
func (i *IT) Administrasi() {}
func (i *IT) Maintenance()  {}

type Akuntan struct{}

func (a *Akuntan) MenghitungData(data int) {}
func (a *Akuntan) Administrasi()           {}
func (a *Akuntan) AnalisaData()            {}

type Mekanik struct{}

func (m *Mekanik) MenghitungData(arr []int) {}
func (m *Mekanik) Maintenance()             {}
func (m *Mekanik) Administrasi()            {}

type Teknisi interface {
	Maintenance()
}

type Analis interface {
	Administrasi()
	MenghitungData(arr []int) int
}

func main() {
	it1 := IT{}
	mk1 := Mekanik{}
	ak1 := Akuntan{}

	var tekni1 Teknisi
	tekni1 = &mk1
	// tekni1 = &it1
	// mk1.
	tekni1.Maintenance()

	// var analis1 Analis
	// analis1 = &it1
	// analis1 = &mk1
	// analis1 = &ak1
}
