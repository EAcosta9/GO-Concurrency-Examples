package main

import (
	"fmt"
)

func funcFirst(D, G chan int) {
	for i := 0; i < 5; i++ {
		D <- i
	}
	G <- -1
}

func funcSecond(E, F chan int) {
	<-F
	E <- 5
}

func main() {
	fmt.Println("Starting Program")
	var C1 chan int = make(chan int)
	var C2 chan int = make(chan int)
	var ChannelShared = make(chan int)
	go funcFirst(C1, ChannelShared)
	go funcSecond(C2, ChannelShared)

LabelA:
	for {
		fmt.Println("Going Through the for loop")
		select {
		case result := <-C1:
			fmt.Printf("This is C1 over here with my value %v \n", result)

		case res2 := <-C2:
			fmt.Printf("This is C2 over here with my value %v \n", res2)
			break LabelA
		}
	}
	fmt.Println("We're at the end of this program. Main is now exiting")
}
