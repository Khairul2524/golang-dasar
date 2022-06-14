package main

import "fmt"

func main() {
	var nilaiujian = 70
	var nillaiabsen = 98

	var lulusujian = nilaiujian > 80
	var lulusabsen = nillaiabsen > 70
	fmt.Println(lulusujian)
	fmt.Println(lulusabsen)
	var lulus = lulusujian && lulusabsen

	fmt.Println(lulus)

	fmt.Println(nilaiujian >= 80 && nillaiabsen >= 70)

}
