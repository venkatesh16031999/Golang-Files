package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// ReadWriteMutex 
var m = sync.RWMutex{}

func main(){
	
	ch := make(chan int,50);
	
	wg.Add(1)
		
	go func(ch <-chan int){
		
		for value := range ch {
			fmt.Println(value);	
		}
		
		wg.Done()
	}(ch)
	
	go func(ch chan<- int){
		ch <- 50
		ch <- 51
		ch <- 52
		close(ch);
	}(ch)
	
	// Read Lock and Unlock
	
	m.RLock();
	m.RUnlock();
	
	// Write Lock and Unlock
	
	m.Lock();
	m.Unlock();
	
	wg.Wait()
}
