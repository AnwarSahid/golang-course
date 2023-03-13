package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			for i, huruf := range "C A M A P B O" {
				if i%2 == 0 {
					fmt.Printf("%#U starts at byte position %d\n", huruf, i)
				}
			}
		} else {
			fmt.Println("Nilai j = ", j)
		}
	}

}
