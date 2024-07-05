package main

import (
	"fmt"
	"strings"
)

func isanagram(word1 string, word2 string) bool {
	// Normalisasi huruf menjadi huruf kecil
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	// Jika panjang kedua kata berbeda, langsung return false
	if len(word1) != len(word2) {
		return false
	}

	// Map untuk menghitung kemunculan huruf
	counts1 := make(map[rune]int)
	counts2 := make(map[rune]int)

	// Hitung kemunculan huruf untuk kata pertama
	for _, char := range word1 {
		counts1[char]++
	}

	// Hitung kemunculan huruf untuk kata kedua
	for _, char := range word2 {
		counts2[char]++
	}

	// Bandingkan jumlah kemunculan huruf
	for char, count := range counts1 {
		if counts2[char] != count {
			return false
		}
	}

	return true
}

func main() {
	// Contoh penggunaan fungsi
	fmt.Println(isanagram("Otto", "Toto"))
	fmt.Println(isanagram("Ayam", "Maya"))
	fmt.Println(isanagram("Tamat", "Kiamat"))
}
