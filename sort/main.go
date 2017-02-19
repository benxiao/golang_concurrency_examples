package main

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
)


//func init(){
//	runtime.GOMAXPROCS(4)
//}


func main() {
	n := 1000000
	runtime.GOMAXPROCS(8)
	array := make([]int, n)
	for i:=0;i!=n;i++{
		array[i]=i
	}


	randomized := make([]int, len(array))
	for i, v := range rand.Perm(len(array)){
		randomized[v] = array[i]
	}

	fmt.Println(randomized)
	start := time.Now()
	quicksort(randomized, 0, len(array))
	duration := time.Since(start)
	fmt.Println(randomized)
	fmt.Println(duration)
}


func quicksort_aux(array []int, s int, e int){
	if e-s > 1 {
		pivot := partition(array, s, e)

		quicksort_aux(array, s, pivot)

		quicksort_aux(array, pivot+1, e)

	}
}

// run twice as fast
func quicksort(array []int, s int, e int){
	if e-s > 1 {
		done := make(chan bool)
		pivot := partition(array, s, e)

		go func() {
			quicksort_aux(array, s, pivot)
			done<-true
		}()

		go func() {
			quicksort_aux(array, pivot+1, e)
			done<-true

		}()
		<-done
		<-done
	}
}


func partition(array []int, s int, e int) int {
	pivot_index := s
	pivot_value := array[s]
	for j:=s+1; j!=e; j++ {
		if array[j] < pivot_value{
			pivot_index += 1
			swap(array, pivot_index, j)
		}
	}
	swap(array, s, pivot_index)
	return pivot_index
}

func swap(array []int, p1 int, p2 int){
	temp := array[p1]
	array[p1] = array[p2]
	array[p2] = temp
}