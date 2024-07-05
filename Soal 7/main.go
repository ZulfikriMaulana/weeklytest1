package main

import (
	"fmt"
)

func sameelement(nums1 []int, nums2 []int) []int {
	// Map untuk menyimpan elemen-elemen yang ada di nums1
	elements := make(map[int]bool)
	// Slice untuk menyimpan hasil akhir
	var result []int

	// Isi map dengan elemen-elemen dari nums1
	for _, num := range nums1 {
		elements[num] = true
	}

	// Iterasi setiap angka dalam nums2
	for _, num := range nums2 {
		// Jika angka ada di map, tambahkan ke hasil dan hapus dari map
		if elements[num] {
			result = append(result, num)
			delete(elements, num)
		}
	}

	return result
}

func main() {
	// Contoh penggunaan fungsi dengan catatan
	nums1a := []int{1, 2, 4, 7, 8}
	nums2a := []int{2, 3, 7, 9}
	result1 := sameelement(nums1a, nums2a)
	if len(result1) == 0 {
		fmt.Println("Output: blank")
	} else {
		fmt.Printf("Output: %v\n", result1)
	}

	nums1b := []int{1, 2, 7, 4, 7, 8}
	nums2b := []int{7, 7, 3, 2, 9}
	result2 := sameelement(nums1b, nums2b)
	if len(result2) == 0 {
		fmt.Println("Output: blank")
	} else {
		fmt.Printf("Output: %v\n", result2)
	}

	nums1c := []int{2, 4, 6, 8}
	nums2c := []int{1, 3, 5, 7, 9}
	result3 := sameelement(nums1c, nums2c)
	if len(result3) == 0 {
		fmt.Println("Output: blank")
	} else {
		fmt.Printf("Output: %v\n", result3)
	}
}
