package main

import (
	"fmt"
)

func main() {
	var person map[string]string = make(map[string]string)

	person["nama"] = "kahirul"
	person["alamat"] = "Beraim"
	person["pekerjaan"] = "Programmer"

	fmt.Println(person)

	fmt.Println(len(person))

	buku := map[string]string{
		"judul":   "Taman Taman Orang Jatuh Cinta dan Memendam Rindu",
		"halaman": "340",
	}
	fmt.Println(buku)

	delete(buku, "halaman")
	fmt.Println(buku)
}
