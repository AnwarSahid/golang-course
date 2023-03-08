package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j < 5; j++ {
		fmt.Println("Nilai j = ", j)
	}
	for j := 0; j <= 12; j += 2 {
		fmt.Println("Character  = ", j)
	}
	for j := 6; j <= 10; j++ {
		fmt.Println("Nilai j  = ", j)
	}
}
