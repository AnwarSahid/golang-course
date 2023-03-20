package main

import (
	"fmt"
	"time"
)

func coba(say int) {
	fmt.Println("Coba ", say)
}
func bisa(say int) {
	fmt.Println("Bisa ", say)

}

func main() {
	for i := 0; i < 4; i++ {
		go coba(i)
		go bisa(i)
	}
	time.Sleep(time.Second * 5)
}
