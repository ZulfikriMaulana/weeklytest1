package main

import (
	"fmt"
)

func plusone(nums []int) []int {
	n := len(nums)

	// Tambahkan 1 ke angka terakhir
	nums[n-1] += 1

	// Tangani carry jika ada
	for i := n - 1; i > 0; i-- {
		if nums[i] == 10 {
			nums[i] = 0
			nums[i-1] += 1
		}
	}

	// Jika digit pertama menjadi 10, perlu menangani carry di tempat pertama
	if nums[0] == 10 {
		nums[0] = 0
		nums = append([]int{1}, nums...)
	}

	return nums
}

func main() {
	// Contoh penggunaan fungsi
	fmt.Println(plusone([]int{1, 2, 3, 4}))
	fmt.Println(plusone([]int{1, 4, 8, 9}))
	fmt.Println(plusone([]int{9, 9, 9, 9}))
}
