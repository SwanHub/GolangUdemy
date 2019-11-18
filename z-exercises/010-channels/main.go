package main

import (
	"fmt"
)

// channels are the pipes that connect concurrent go routines
// you can send a value into a channel from one go routine, and pull it down from another

func main() {
	exercise1a() // basic channel creation, avoid blocking with go routine
	exercise1b() // basic channel creation, avoid blocking with buffering
	exercise2()  // basic send, receive channel
	exercise3()  // directional channels
	exercise4()  // using the select statement
	exercise5()  // comma ok
	exercise6()  // write and read from buffered channel
	exercise7()  // write and read multiple loops, 10 go routines
}

func exercise1a() {
	fmt.Println("------exercise 1------")
	c := make(chan string)
	go func() {
		c <- "the meaning of life"
	}()
	fmt.Println(<-c)
}

func exercise1b() {
	fmt.Println("------exercise 2------")
	c := make(chan string, 1)
	c <- "the meaning of life"
	fmt.Println(<-c)
}

func exercise2() {
	fmt.Println("------exercise 2------")
	cs := make(chan int)
	go func() {
		cs <- 42222
	}()

	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func exercise3() { // understanding sending and receiving from and to channels -- good challenge
	fmt.Println("------exercise 3------")

	c := make(chan int)
	// feels like there should be a go routine going on here.
	go send3(c)

	receive3(c)

	fmt.Println("about to exit")
}

func send3(cs chan int) <-chan int {
	for i := 0; i < 100; i++ {
		cs <- i
	}
	close(cs)
	return cs
}

func receive3(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func exercise4() {
	fmt.Println("------exercise 4------")
	q := make(chan int) // first channel
	c := send4(q)

	receive4(c, q)

	fmt.Println("about to exit")
}

func send4(q chan<- int) <-chan int {
	c := make(chan int) // second channel

	go func() {
		for i := 0; i < 50; i++ {
			c <- i
		}
		q <- 50
	}()

	return c
}

func receive4(c, q <-chan int) {
	for { // infinitely listen, until broken
		select {
		case msg1 := <-c:
			fmt.Println(msg1)
		case <-q:
			return
		}
	}
}

func exercise5() {
	fmt.Println("------exercise 5------")
	c := make(chan int)
	go func() {
		c <- 1
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	close(c)
	v, ok = <-c
	fmt.Println(v, ok)
}

func exercise6() {
	fmt.Println("------exercise 6------")
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func exercise7() {
	fmt.Println("------exercise 7-------")

	for i := 0; i < 10; i++ {
		fmt.Println("round ", i+1)
		c := make(chan int)
		go write10(c)
		read10(c)
	}
}

func write10(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func read10(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
