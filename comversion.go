package main

import "fmt"

func main() {
	var int32 = 10000
	var int64 = int64(int32)
	var int8 = int8(int32)

	fmt.Println(int32)
	fmt.Println(int64)
	fmt.Println(int8)

	var Nama = "Khairul Kahpi"
	var e byte = Nama[0]
	var eString string = string(e)

	fmt.Println(Nama)
	fmt.Println(e)
	fmt.Println(eString)
}
