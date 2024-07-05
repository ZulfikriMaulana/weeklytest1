package main

import (
	"fmt"
)

func removeduplicate(nums []int) []int {
	// Map untuk menyimpan angka yang sudah muncul
	seen := make(map[int]bool)
	// Slice untuk menyimpan hasil akhir
	var result []int

	// Iterasi setiap angka dalam slice
	for _, num := range nums {
		// Jika angka belum muncul sebelumnya, tambahkan ke hasil
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func main() {
	// Contoh penggunaan fungsi
	nums := []int{1, 1, 2, 2, 3, 3, 4, 4, 4}
	fmt.Println(removeduplicate(nums)) // Output: [1, 2, 3, 4]
}
