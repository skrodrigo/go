package main

import (
	"fmt"
	"time"
)

// Theard 1
func Channel() {

 // Theard 1 <-> Thread 2
	hello := make(chan string)

 // Thread 2
	go func() {
		hello <- "Hello World"
	}()

	result := <-hello
	fmt.Println(result)
}

func Worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker", workerId, "processed", res)
    time.Sleep(time.Second)	
	}
}

func WebServer(){
	msg := make(chan int)
	go Worker(1, msg)
	go Worker(2, msg)
	for i := 0; i < 10; i++ {
		msg <- i
	}
}