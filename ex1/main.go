package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i:=0; i<5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value
	}
	fmt.Println("You're boring; I am leaving.")
}

func boring(msg string, c chan string){
	for i:=0; ;i++{
		c<-fmt.Sprintf("%v %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
