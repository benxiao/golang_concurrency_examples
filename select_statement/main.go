package main

import (
	"fmt"
	"time"
)

func main() {
	joe := boring("Joe", 300)
	ann := boring("Ann", 100)
	fennel := fanIn(joe, ann)
	for i := 0; i < 100; i++ {
		fmt.Printf("You say: %q\n", <-fennel) // force joe and ann to run in locked steps
	}
	fmt.Println("You're boring; I am leaving.")
}

func boring(msg string, t int) <-chan string { //Return receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ { // We launch the goroutine from inside the function
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(t) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	for {
		select{
		case s1:= <-input1:
			c<-s1
		case s2:= <-input2:
			c<-s2
		}
	}
	return c

}
