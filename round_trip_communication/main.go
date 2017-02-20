package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i:= 0;i<10;i++{fmt.Println(<-c)}
	quit<-"Bye" // kick off the case <-quit and subsquent cleanups
	fmt.Printf("Joe says: %q\n", <-quit) //wait on Joe
}


func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func(){
		for i:=0;;i++{
			select{
			case c <- fmt.Sprintf("%s %d", msg, i):
				//pass
			case <- quit:
				cleanup()
				quit<-"See you!"
				return
			}
		}
	}()
	return c
}

func cleanup(){
	time.Sleep(100*time.Millisecond)
	fmt.Println("cleanup has completed!")
}



