//WaitGroup

package main

import (
	"fmt"
	"sync"
	"time"
)

func goRoutine1(wg *sync.WaitGroup) {
	time.Sleep(5000 * time.Millisecond)
	wg.Done()
	fmt.Println("go routine 1 completed")
}

func goRoutine2(wg *sync.WaitGroup) {
	time.Sleep(2000 * time.Millisecond)
	wg.Done()
	fmt.Println("go routine 2 completed")
}

func goRoutine3(wg *sync.WaitGroup) {
	time.Sleep(1000 * time.Millisecond)
	wg.Done()
	fmt.Println("go routine 3 completed")
}

func main() {
	wg := &sync.WaitGroup{}
	/*
		wg.Add(1)
		go goRoutine1(wg)
		wg.Add(1)
		go goRoutine2(wg)
		wg.Add(1)
		go goRoutine3(wg)
	*/
	wg.Add(3)
	go goRoutine1(wg)
	go goRoutine2(wg)
	go goRoutine3(wg)
	wg.Wait()
	fmt.Println("main completed")
}
