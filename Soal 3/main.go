package main

import (
	"fmt"
	"strings"
)

func capitalize(words []string, except []string) []string {
	// Buat map untuk kata-kata yang tidak ingin dikapitalisasi
	exceptMap := make(map[string]bool)
	for _, word := range except {
		exceptMap[word] = true
	}

	// Slice untuk hasil akhir
	var result []string

	// Iterasi melalui setiap kata dalam slice kata-kata
	for _, word := range words {
		// Jika kata ada pengecualian, tambahkan tanpa perubahan
		if exceptMap[word] {
			result = append(result, word)
		} else {
			// Jika tidak, kapitalisasi huruf pertama dan tambahkan ke hasil
			capitalizedWord := strings.Title(word)
			result = append(result, capitalizedWord)
		}
	}

	return result
}

func main() {
	// Contoh penggunaan fungsi
	words := []string{"this", "is", "a", "konoha"}
	except := []string{"is", "a"}
	fmt.Println(capitalize(words, except)) // Output: ["This", "is", "a", "Konoha"]
}
