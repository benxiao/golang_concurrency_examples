package main

import (
	"fmt"
	"time"
)



func main() {
	c := fanIn(boring("Ann", 300), boring("Joe", 100))
	for i:=0; i < 5;i++{
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait<-true
		msg2.wait<-true
	}
}


type Message struct {
	str string
	wait chan bool
}


func boring(msg string, t int) <-chan Message { //Return receive-only channel of strings
	c := make(chan Message)
	go func() {
		wait := make(chan bool)
		for i := 0; ; i++ { // We launch the goroutine from inside the function
			c <- Message{fmt.Sprintf("%s %d", msg, i), wait}
			time.Sleep(time.Duration(t) * time.Millisecond)
			<-wait
		}

	}()
	return c // Return the channel to the caller
}


func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func(){
		for{
			c<-<-input1
		}
	}()

	go func(){
		for{
			c<-<-input2
		}
	}()
	return c

}