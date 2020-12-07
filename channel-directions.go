package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg

	//수신 시도시 에러
	//test := <-pings
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
