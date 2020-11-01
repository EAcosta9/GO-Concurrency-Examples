package main

import (
	"fmt"
	"time"
)

func funcFirst(D chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		D <- i
	}
}

func funcSecond(E chan int) {
	time.Sleep(6 * time.Second)
	E <- 5
}

func main() {
	fmt.Println("Starting Program")
	var C1 chan int = make(chan int)
	var C2 chan int = make(chan int)

	go funcFirst(C1)
	go funcSecond(C2)

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
