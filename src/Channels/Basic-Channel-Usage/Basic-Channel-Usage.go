package main

import (
	"fmt"
	"time"
)

func Launcher1(C chan int) {
	fmt.Println("I'll send something over the channel in about five seconds")
	time.Sleep(5 * time.Second)
	C <- 5
}

func main() {
	fmt.Println("Starting Program")
	Chan1 := make(chan int)
	go Launcher1(Chan1)

	fmt.Println("The receive from the channel will now be blocking")
	num := <-Chan1
	fmt.Printf("Received value %v from the channel\n", num)

	fmt.Println("Now exiting program")
}
