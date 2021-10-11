package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var ch = make(chan int, 1)
var wg sync.WaitGroup


func worker(id int) {
	ch <- 1
	counter ++
	<- ch
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker(0)
	}
	wg.Wait()
	fmt.Println(counter )


}
