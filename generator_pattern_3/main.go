package main

import (
	"fmt"
	//"math/rand"
	"time"
)

func main() {
	joe := boring("Joe", 300)
	ann := boring("Ann", 100)
	fennel := fanIn2(joe, ann)
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

func fanIn2(inputs ...<-chan string) <-chan string {
	c := make(chan string)
	fmt.Println(inputs)
	for i, input := range inputs{
		fmt.Printf("%d %T\n",i, input)
		go func(input <-chan string){
			for{c<-<-input}
		}(input)
	}
	return c
}
