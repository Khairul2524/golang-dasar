package main

import "fmt"

func main() {
	var nilai [3]string
	nilai[0] = "Khairul"
	nilai[1] = "Kahpi"
	nilai[2] = "Al Mubarrok"

	fmt.Println(nilai[0])
	fmt.Println(nilai[1])
	fmt.Println(nilai[2])

	var values = [3]int{
		80,
		80,
		70,
	}
	fmt.Println(values)

	var baru [10]int
	fmt.Println(len(baru))
}
