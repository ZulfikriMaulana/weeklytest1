package main

import (
	"fmt"
)

func countFrequentElement(nums []int) map[int]int {
	frequency := make(map[int]int)

	// Menghitung frekuensi kemunculan setiap elemen dalam slice
	for _, num := range nums {
		frequency[num]++
	}

	return frequency
}

func main() {
	// Contoh penggunaan fungsi
	nums1 := []int{1, 2, 3, 4, 4, 4, 3, 3, 2, 4}
	fmt.Println("Input:", nums1)
	fmt.Println("Output:", countFrequentElement(nums1))

	nums2 := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	fmt.Println("Input:", nums2)
	fmt.Println("Output:", countFrequentElement(nums2))
}
