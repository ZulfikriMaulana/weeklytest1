package main

import (
	"fmt"
)

func oddbeforeeven(nums []int) []int {
	// Slice untuk menyimpan angka ganjil dan genap
	var odds []int
	var evens []int

	// Iterasi setiap angka dalam slice
	for _, num := range nums {
		if num%2 != 0 {
			odds = append(odds, num)
		} else {
			evens = append(evens, num)
		}
	}

	// Gabungkan slice ganjil dan genap
	return append(odds, evens...)
}

func main() {
	// Contoh penggunaan fungsi
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(oddbeforeeven(nums))
}
