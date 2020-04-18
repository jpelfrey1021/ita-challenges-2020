package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(helloWorld())
}

func helloWorld() string {
	ch1 := make(chan string)

	go hello(ch1)
	go world(ch1)

	st := <-ch1 + " " + <-ch1
	return st
}

func hello(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "hello"
}

func world(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "world"
}
