package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := make(chan string), make(chan string)
	go func(){a<-"a"}()
	go func(){b<-"b"}()
	if (rand.Intn(2)==0){
		a = nil
		fmt.Println("a is blocked")
	}else{
		b = nil
		fmt.Println("b is blocked")
	}
	select{
	case v1:=<-a:
		fmt.Println("select",v1)
	case v2:=<-b:
		fmt.Println("select",v2)
	}

}
