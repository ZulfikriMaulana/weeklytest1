package main

import (
	"fmt"
)

func removeduplicateletter(word string) string {
	// Map untuk menyimpan huruf yang sudah muncul
	seen := make(map[rune]bool)
	// StringBuilder untuk menyusun hasil akhir
	var result []rune

	// Iterasi setiap huruf dalam kata
	for _, char := range word {
		// Jika huruf belum muncul sebelumnya, tambahkan ke hasil
		if !seen[char] {
			seen[char] = true
			result = append(result, char)
		}
	}

	// Konversi slice of rune menjadi string dan kembalikan
	return string(result)
}

func main() {
	// Contoh penggunaan fungsi
	fmt.Println(removeduplicateletter("bananasna"))
	fmt.Println(removeduplicateletter("lalalamama"))
}
