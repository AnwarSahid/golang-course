package main

import (
	"fmt"
	"strconv"
)

func main() {

	word := "selamat malam"
	freq := make(map[rune]int)

	for _, c := range word {
		freq[c]++
	}
	hasil := "map["
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c \n", word[i])

	}
	for c, count := range freq {
		huruf := string(c)
		jumlah := strconv.Itoa(count)
		hasil += huruf + ":" + jumlah + ", "
	}
	hasil += "]"

	fmt.Println(hasil)
}
