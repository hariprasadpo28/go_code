package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var c = make(chan int)
var wg = sync.WaitGroup{}
func rating() {
	var r = rand.Intn(5)
	c <- r
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	wg.Done()
}

func main(){
	var sum int = 0
	for i:=0; i<200; i++{
		wg.Add(1)
		go rating()
		sum += <-c
	}
	fmt.Println("Average rating: ", sum/200)
	wg.Wait()

}