package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	exercise1() // launch basic go routines
	exercise2() // method sets, interfaces
	// exercise3() // using go routines to create a race condition (not good)
	exercise4() // use mutex to avoid data race
	exercise5() // use sync/atomic package to avoid data race // note: done out of order
	exercise6() // printing runtime info (OS & ARCH) to the screen
}

func exercise1() {
	fmt.Println("------exercise 1------")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("go routine one")
		wg.Done()
	}()

	go func() {
		fmt.Println("go routine 2")
		wg.Done()
	}()
	wg.Wait()
}

func exercise2() {
	fmt.Println("------exercise 2-------")
	p1 := person{"josh", "withers", 14}
	// saySomething(p1) can't do this because speak only takes a pointer
	saySomething(&p1)
}

type person struct {
	First string
	Last  string
	Age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("My name is", p.First, p.Last)
	fmt.Println("\tI'm this many years old: ", p.Age)
}

func saySomething(h human) {
	h.speak()
}

// incrementor program
func exercise3() {
	fmt.Println("------exercise 3------")
	counter := 0

	var wg sync.WaitGroup
	gs := 10
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			fmt.Println("GoRoutine # ", counter)
			v := counter
			runtime.Gosched() // yield to the processor // not needed when we do mutex
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter end: ", counter)
	// $go run -race main.go
}

func exercise4() {
	fmt.Println("------exercise 4------")
	counter := 0

	var wg sync.WaitGroup
	var mu sync.Mutex
	gs := 10
	wg.Add(gs) // adds 10 go routines

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			fmt.Println("GoRoutine # ", counter)
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter end: ", counter)
	// $go run -race main.go
}

func exercise5() {
	fmt.Println("------exercise 5------")

	var counter int64
	var wg sync.WaitGroup
	gs := 10
	wg.Add(gs) // adds 10 go routines

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter end: ", counter)
}

func exercise6() {
	fmt.Println("------exercise 6------")
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("GO ARCH: ", runtime.GOARCH)
}
