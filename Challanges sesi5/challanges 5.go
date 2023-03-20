package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	coba := make([]interface{}, 4)
	coba[0] = []string{"Coba1", "Coba2", "Coba3"}
	coba[1] = []string{"Coba1", "Coba2", "Coba3"}
	coba[2] = []string{"Coba1", "Coba2", "Coba3"}
	coba[3] = []string{"Coba1", "Coba2", "Coba3"}
	bisa := make([]interface{}, 4)
	bisa[0] = []string{"bisa1", "bisa2", "bisa3"}
	bisa[1] = []string{"bisa1", "bisa2", "bisa3"}
	bisa[2] = []string{"bisa1", "bisa2", "bisa3"}
	bisa[3] = []string{"bisa1", "bisa2", "bisa3"}

	for i := 0; i < 4; i++ {
		wg.Add(1)

		go func(idx int) {
			mu.Lock()
			defer mu.Unlock()

			fmt.Printf("Coba %d: %v\n", idx+1, coba[idx])

			wg.Done()
		}(i)
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)

		go func(var2 int) {
			mu.Lock()
			defer mu.Unlock()

			fmt.Printf("Bisa %d: %v\n", var2+1, bisa[var2])

			wg.Done()
		}(i)
	}

	wg.Wait()
}
