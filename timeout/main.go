package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	quit := time.After(10*time.Second)
	c := boring("Joe")
	for {
		select {
		case s:=<-c:
			fmt.Println(s)
		case <-time.After(1*time.Second): // used to time out a message
			fmt.Println("You are too slow.")
		case <-quit: // used to end the conversation
			fmt.Println("We are done!")
			return
		}
	}
}

func boring(msg string) <-chan string{
	c := make(chan string)
	go func(){
		for i:=0; ;i++{
			c<-fmt.Sprintf("%v %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()
	return c
}