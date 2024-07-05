package main

import (
	"fmt"
)

func addDigits(nums1 []int, nums2 []int) []int {
	len1 := len(nums1)
	len2 := len(nums2)
	maxLen := max(len1, len2)

	// Buat slice untuk menyimpan hasil penjumlahan
	sum := make([]int, maxLen)

	// Penambahan digit per digit
	carry := 0
	for i := 0; i < maxLen; i++ {
		digit1 := 0
		if i < len1 {
			digit1 = nums1[len1-1-i]
		}

		digit2 := 0
		if i < len2 {
			digit2 = nums2[len2-1-i]
		}

		digitSum := digit1 + digit2 + carry
		sum[maxLen-1-i] = digitSum % 10
		carry = digitSum / 10
	}

	// Jika masih ada carry setelah iterasi, tambahkan ke hasil
	if carry > 0 {
		sum = append([]int{carry}, sum...)
	}

	return sum
}

// Fungsi helper untuk mencari maksimum dari dua bilangan
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Contoh penggunaan fungsi
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}
	fmt.Println("Input:", nums1, nums2)
	fmt.Println("Output:", addDigits(nums1, nums2))

	nums3 := []int{9, 2, 7}
	nums4 := []int{1, 3, 5}
	fmt.Println("Input:", nums3, nums4)
	fmt.Println("Output:", addDigits(nums3, nums4))
}
