package main

import (
	"fmt"
	"time"
)

func Launcher1(C chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		C <- i
	}
}

func main() {
	fmt.Println("Starting Program")
	Chan1 := make(chan int)
	timer := time.NewTimer(5 * time.Second)

	go Launcher1(Chan1)

LabelA:
	for {
		select {
		case <-timer.C:
			fmt.Println("Timer went off because nothing was received on Chan1. I am breaking out of this for loop.")
			break LabelA

		case num := <-Chan1:
			fmt.Printf("From Chan1 I have received %v\n", num)
			fmt.Println("I will now reset the timer")
			timer.Stop()
			timer.Reset(5 * time.Second)

		}
	}
	fmt.Println("I have hit the end of main. Now exiting the program")
}
