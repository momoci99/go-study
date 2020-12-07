package main

import "fmt"

func main() {

	//여기서도 make가..?!
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

//보충
/* http://golang.site/go/article/22-Go-%EC%B1%84%EB%84%90 */
