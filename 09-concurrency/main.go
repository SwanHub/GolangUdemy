package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Concurrency is a DESIGN PATTERN that allows you to run code in parallel IF you have dual core
// Parallelism is the EXECUTION or the DOING of concurrent code
// Concurrency is about COMPOSING code to deal with multiple processes at once
func main() {
	waitGroup()                   // using sync.WaitGroup for .Add() .Wait()
	goRoutinesAndRaceConditions() // setting up 100 go routines and seeing race condition // errors
	mutex()                       // locking down our work, avoids data races
	atomicFunction()              // another way to lock down your code from data races
}

var wg sync.WaitGroup

// similarities between async fetching and goroutines on multiple cores
func waitGroup() {
	fmt.Println("------using waitgroup, setting up basic go routine------")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go first()
	second()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func first() {
	for i := 0; i < 5; i++ {
		fmt.Println("first:", i)
	}
	wg.Done()
}

func second() {
	for i := 0; i < 5; i++ {
		fmt.Println("second:", i)
	}
}

// multiplexer == mux
// many into one

func goRoutinesAndRaceConditions() {
	fmt.Println("------using waitgroup, experiencing data races------")
	fmt.Println("CPU # ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	counter := 0

	gors := 100 // 100 go routines
	var wg sync.WaitGroup
	wg.Add(gors)

	for i := 0; i < gors; i++ {
		go func() {
			fmt.Println("GoRoutine # ", counter)
			v := counter
			// time.Sleep(time.Second) // another way to work w async, similar to "setTimeout"
			runtime.Gosched() // run something else if you want to run something else
			v++
			counter = v
			wg.Done() // close the go routine at this point
		}()
	}
	wg.Wait() // don't exist program until everything is done

	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
	// $go run -race main.go
}

// this code isn't doing what it's supposed to do. Data races still occurring.
func mutex() {
	fmt.Println("------using mutex and waitgroup, preventing data races------")
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

func atomicFunction() {
	fmt.Println("------using waitgroup and atomic, preventing data races------")
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64 // int64 is good indication that we're using package atomic

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
