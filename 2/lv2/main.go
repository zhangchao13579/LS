package main

import "fmt"

var Wallet = 0

func main() {
	ch := make(chan int, 10000)
	for i := 0; i < 10000; i++ {
		ch <- 50
		Wallet += <-ch
	}
	close(ch)
	fmt.Println("泡泡现在有", Wallet, "元")
}
