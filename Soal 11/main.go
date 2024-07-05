package main

import (
	"fmt"
)

func sliceFruits(fruits1 []string, fruits2 []string, operationType int) []string {
	switch operationType {
	case 1:
		// Same: Buah yang sama di kedua slice
		return intersectionString(fruits1, fruits2)
	case 2:
		// Difference: Buah yang ada di fruits1 tetapi tidak ada di fruits2
		return differenceString(fruits1, fruits2)
	default:
		// Jika operationType tidak valid, kembalikan slice kosong
		return []string{}
	}
}

// Helper function untuk mencari Intersection dari dua slice string
func intersectionString(slice1 []string, slice2 []string) []string {
	set := make(map[string]bool)
	var result []string

	for _, item := range slice1 {
		set[item] = true
	}

	for _, item := range slice2 {
		if set[item] {
			result = append(result, item)
			delete(set, item)
		}
	}

	return result
}

// Helper function untuk mencari Difference dari dua slice string
func differenceString(slice1 []string, slice2 []string) []string {
	set := make(map[string]bool)
	var result []string

	for _, item := range slice2 {
		set[item] = true
	}

	for _, item := range slice1 {
		if !set[item] {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	// Contoh penggunaan fungsi
	fruits1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	fruits2 := []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}

	// Operasi Same (operationType = 1)
	fmt.Println("Operation Type = 1 (Same):", sliceFruits(fruits1, fruits2, 1))

	// Operasi Difference (operationType = 2)
	fmt.Println("Operation Type = 2 (Difference):", sliceFruits(fruits1, fruits2, 2))
}
