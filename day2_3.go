package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var bal int = 500
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}

func deposit(v int){
	m.Lock()
	bal += v
	fmt.Println("Deposited ", v, "\tCurrent Balance: ",bal)
	m.Unlock()
	wg.Done()

}

func withdraw(v int){
	m.Lock()
	if bal < v {
		fmt.Println("Insufficient Balance for withdrawing ",v,"\tCurrent Balance: ", bal)
		m.Unlock()
		wg.Done()
		return
	} else{
		bal -= v
		fmt.Println("Debited ",v,"\tCurrent Balance: ", bal)

	}
	m.Unlock()
	wg.Done()
}


func main() {
	var amount int
	for i := 0;i<5;i++{
		amount = rand.Intn(500)
		wg.Add(2)
		go deposit(amount)
		amount = rand.Intn(1000)
		go withdraw(amount)
	}
	wg.Wait()

}


