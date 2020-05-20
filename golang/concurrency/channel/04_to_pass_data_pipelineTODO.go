// example of channels to pass the data from one go routine to another go routine: pipelines
// pipeline scenario: pizza delivery service
// go routines: 
// 		1) buy flour
// 		2) cut veg and prepare dough
//		3) deliver


package main

import (
	"fmt"
)


func main() {
	ch:=make(chan int)
	outCh:=make(chan int)
	data := []int{2,3,4,5}
	go stage1(2,3,ch)
	go stage2(3,ch,outCh)
	fmt.Println(<-outCh)
	
}

func buy(done chan<- bool) {
	ch<-x+y
}

func prepare(y int, ch <-chan int, out chan<- int)  {
	x:=<-ch
}

 func deliver ///