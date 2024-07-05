package main

import (
	"fmt"
)

func sliceOperation(nums1 []int, nums2 []int, typeOperation int) []int {
	switch typeOperation {
	case 1:
		// Difference1 (A - B)
		return difference(nums1, nums2)
	case 2:
		// Difference2 (B - A)
		return difference(nums2, nums1)
	case 3:
		// Union (A U B)
		return union(nums1, nums2)
	case 4:
		// Intersection
		return intersection(nums1, nums2)
	default:
		// Jika tipe operasi tidak valid, kembalikan slice kosong
		return []int{}
	}
}

// Helper function untuk mencari Difference (A - B)
func difference(nums1 []int, nums2 []int) []int {
	elements := make(map[int]bool)
	for _, num := range nums2 {
		elements[num] = true
	}

	var result []int
	for _, num := range nums1 {
		if !elements[num] {
			result = append(result, num)
		}
	}

	return result
}

// Helper function untuk mencari Union (A U B)
func union(nums1 []int, nums2 []int) []int {
	elements := make(map[int]bool)
	var result []int

	for _, num := range nums1 {
		if !elements[num] {
			elements[num] = true
			result = append(result, num)
		}
	}

	for _, num := range nums2 {
		if !elements[num] {
			elements[num] = true
			result = append(result, num)
		}
	}

	return result
}

// Helper function untuk mencari Intersection
func intersection(nums1 []int, nums2 []int) []int {
	elements := make(map[int]bool)
	var result []int

	for _, num := range nums1 {
		elements[num] = true
	}

	for _, num := range nums2 {
		if elements[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{2, 4, 6, 7}

	// Difference1 (A - B)
	diff1 := sliceOperation(nums1, nums2, 1)
	fmt.Println("Difference1 (A - B):", diff1)

	// Difference2 (B - A)
	diff2 := sliceOperation(nums1, nums2, 2)
	fmt.Println("Difference2 (B - A):", diff2)

	// Union (A U B)
	uni := sliceOperation(nums1, nums2, 3)
	fmt.Println("Union (A U B):", uni)

	// Intersection
	inter := sliceOperation(nums1, nums2, 4)
	fmt.Println("Intersection:", inter)
}
