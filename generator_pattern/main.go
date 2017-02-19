package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c:=boring("boring!")
	for i:=0; i<5;i++{
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I am leaving.")
}


func boring(msg string)<-chan string { //Return receive-only channel of strings
	c := make(chan string)
	go func(){
		for i:=0;;i++{ // We launch the goroutine from inside the function
			c<-fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller
}