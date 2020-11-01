package main

import (
	"fmt"
)

func Launcher2(C2 chan int) {
	C2 <- -1
}

func Launcher1(C , D chan int) {
	<-D
	C <- 5
}

func main() {
	fmt.Println("Starting Program")
	Chan1 := make(chan int)
	Chan2 := make(chan int)
	
	go Launcher1(Chan1,Chan2)
	go Launcher2(Chan2)

	fmt.Println("The receive from the channel will now be blocking")
	num := <-Chan1
	fmt.Printf("Received value %v from the channel\n", num)

	fmt.Println("Now exiting program")
}
