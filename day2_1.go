package main

import (
	"fmt"
	"sync"
)

var m=sync.Mutex{}
var wg = sync.WaitGroup{}
var f = make(map[string]int, 26)

func checkFrequency(word string){
	m.Lock()
	for i:=0; i<len(word);i++{
		f[string(word[i])] += 1
	}
	m.Unlock()
	wg.Done()
}

func main(){
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	//allFreq := make(map[string]map[string]int,4)
	for i:= range words{
		wg.Add(1)
		go checkFrequency(words[i])
	}
	wg.Wait()
	fmt.Println(f)
}