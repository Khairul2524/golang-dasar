package main

import "fmt"

func main() {
	var nama = "Lukman"
	if nama == "Kahfie" {
		fmt.Println("hello word")
	} else if nama == "Robby" {
		fmt.Println("Hello Robby")
	} else if nama == "Lukman" {
		fmt.Println("Hello Lukman")
	} else if nama == "ayu" {
		fmt.Println("Hello Ayu")
	} else {
		fmt.Println("Hello", nama)
	}

	if panjang := len(nama); panjang > 5 {
		fmt.Println("nama kepanjangan")
	} else {
		fmt.Println("Nama Kependekan")
	}

}
