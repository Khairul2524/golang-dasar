package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	var k = 10
	k += 10
	fmt.Println(k)

	k++
	fmt.Println(k)

	var negative = -100
	var positive = 100

	fmt.Println(negative)
	fmt.Println(positive)
}
