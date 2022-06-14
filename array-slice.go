package main

import "fmt"

func main() {
	var bulan = [...]string{
		"Januari",
		"Februasi",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	fmt.Println(bulan)

	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// bulan[6] = "juli di ubah"
	// fmt.Println(bulan)

	// slice1[2] = "diubah"
	// fmt.Println(bulan)

	var hari = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Ahad",
	}
	harislice := hari[5:]
	harislice[0] = "Hari Senin"
	harislice[1] = "Hari Selasa"

	fmt.Println(hari)

	harislice2 := append(harislice, "Hari Libur")
	fmt.Println(harislice2)
	harislice2[0] = "Hari Jumat"
	harislice2[1] = "Hari Sabtu"
	fmt.Println(harislice2)

	slicebaru := make([]string, 2, 5)
	slicebaru[0] = "Khairul"
	slicebaru[1] = "Kahpi"
	fmt.Println(slicebaru)
	fmt.Println(len(slicebaru))
	fmt.Println(cap(slicebaru))

	copyslice := make([]string, len(slicebaru), cap(slicebaru))
	copy(copyslice, slicebaru)
	fmt.Println(copyslice)
}
