package main

import "fmt"

type User struct {
	ID   int
	Nama string
	Umur int
}

func (u *User) TambahUmur() {
	u.Umur++
}

func (u User) NaikUmur() {
	u.Umur++
}

func main() {
	user1 := User{1, "Jerry", 26}
	fmt.Println("Kondisi awal:", user1)
	user1.NaikUmur()
	fmt.Println(user1)
	user1.TambahUmur()
	fmt.Println(user1)
}
