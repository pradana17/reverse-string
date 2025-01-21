package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan sebuah string: ")
	input, _ := reader.ReadString('\n')

	// Menghapus karakter newline atau spasi di akhir input
	input = strings.TrimSpace(input)

	reversed := reverseString(input)
	fmt.Println("String asli: ", input)
	fmt.Println("String terbalik: ", reversed)
}
