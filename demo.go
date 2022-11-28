package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Printf(s)
	}

}
func say1(s string, c chan int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf(s)
	}
	c <- 0
	close(c)
}
func main() {
	c := make(chan int)
	go say1("AAA\n", c)
	say("BBB\n")

	<-c
}
