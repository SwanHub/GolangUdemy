package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// channels allow us to pass values between go routines
// channels BLOCK
// it's like an array that has a strict capacity type
// For buffered channels (defined quantity), you don't need to spin upm multiple go routines
// It seems like you generally start out creating ONE bi-directional channel, and then
// you split the send and receive channel functions up into multiple go routines.
// specificity of send / receive increases type safety of the program
// SELECT allows us to wait on / serve multiple go channel operations

func main() {
	// channelBlock() // does not work
	channelThrough() // create channel, input integer via go routine, remove value from channel
	// unsuccessfulBuffer() // exceeds channel capacity, doesn't work
	successfulBuffer()         //create channel with capacity of 1 int, input int, remove int and print
	successfulBuffer2()        // capacity of 2 ints, channel holds 2 ints successfully
	directionalChannel()       // send or receive, split, from / to a channel
	channelSendReceive()       // send and receive-only channels in action w go routine
	rangeOverAndCloseChannel() // using range to print channel values, then closing channel
	selectStatement()          // using select statement to capture values from an active, open channel
	commaOK()                  // using the health checker, comma OK idiom
	moreCommaOK()              // comma ok directly on channel
	fanIn()                    // writing from multiple channels to one channel, fanin design pattern
	fanOut()                   // for an unknown amount of work, creates go routines, then fans it in
	contextOverview()          // advanced topic, overview of contex
}

func channelBlock() {
	// doesn't work, channel is blocked, not yet sure why
	c := make(chan int)
	c <- 4444
	fmt.Println(<-c)
}

func channelThrough() {
	fmt.Println("------successful channel creation-------")
	c := make(chan int)

	go func() { // I don't get why this is necessary (you need to create a go routine)
		// that's the purpose of channels, to communicate between go routines
		c <- 44443
	}()

	fmt.Println(<-c)
}

// the "buffer" part ostensibly means you are determining the quanity of values which the channel
// can hold... otherwise, the channel can just hold one value and to update it, the value must first
// be removed, a la the above
func unsuccessfulBuffer() {
	c := make(chan int, 1)

	c <- 4333
	c <- 443

	fmt.Println(<-c)
}

func successfulBuffer() {
	fmt.Println("------successful channel buffer-------")
	c := make(chan int, 1)

	c <- 33445566

	fmt.Println(<-c)
}

func successfulBuffer2() {
	fmt.Println("------successful channel buffer w multiple values in channel-------")
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}

// there are channels which can only be sent or received to... the ones above were
// BI-DIRECTIONAL channels
func directionalChannel() {
	fmt.Println("------directional channels------")
	c := make(chan int, 2)
	c <- 4099
	c <- 4010

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c) // chan int
	fmt.Println("-------")
	// send channel
	s := make(chan<- int, 2) // send only channel with buffer size of 2
	s <- 3
	s <- 2
	fmt.Printf("Type of send-only chan: %T\n", s)
	// receive channel
	fmt.Println("-------")
	r := make(<-chan int, 2) // send only channel with buffer size of 2
	// <-r CAN do this, but there are no values in the buffered channel at start
	// r <- 3 can't do this, receive-only
	// r <- 2 can't do this, receive-only
	fmt.Printf("Type of receive-only chan: %T\n", r)
	fmt.Println("-------")
	// assigning a LARGER construct to a SMALLER construct is possible
	// assning a SMALLER construct to a LARGER construct is impossible
	fmt.Println("converting from general to specific")
	fmt.Printf("c to \t%T\n", (chan<- int)(c))
	fmt.Printf("%T\n", c) // but the conversion doesn't hold?
	// fmt.Printf("c to \t%T\n", (<-chan int)(s)) doesn't work because specific to specific
}

func channelSendReceive() {
	fmt.Println("------using channels------")
	c := make(chan int)

	// send
	go sendHelper(c)

	//receive
	receiveHelper(c)

	fmt.Println("about to exit")
}

// send
func sendHelper(c chan<- int) {
	c <- 1
	fmt.Println("write 1 to the channel")
}

//receive
func receiveHelper(c <-chan int) {
	fmt.Println(<-c)
}

func rangeOverAndCloseChannel() {
	fmt.Println("------range over channel and close it------")
	c := make(chan int)

	// sending to channel
	go func() { // is a channel an array holding type interface?
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // without this value, will be stuck in terminal looking for values forever
	}()

	// range will loop over channel until it's closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func selectStatement() {
	fmt.Println("------select statement------")
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// takes in three channels, all of type "send-only"
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

// select statement good for pulling values OFF of a channel
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel: ", v)
		case v := <-o:
			fmt.Println("From the odd channel: ", v)
		case v := <-q:
			fmt.Println("From the quit channel: ", v)
			return
		}
	}
}

func commaOK() {
	fmt.Println("----using comma ok idiom to check channels------")
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send2(even, odd, quit)
	receive2(even, odd, quit)
	fmt.Println("about to exit")
}

func send2(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive2(e, o <-chan int, q <-chan bool) {
	// infinite loop until we return out with a "quit" channel reception
	for {
		select {
		case r := <-e:
			fmt.Println("even: ", r)
		case r := <-o:
			fmt.Println("odd: ", r)
		case i, ok := <-q:
			if !ok { // false because zero value
				fmt.Println("from comma ok", i, ok)
				return
			}
			fmt.Println("from comma ok", i)
		}
	}
}

func moreCommaOK() {
	fmt.Println("------more channel work------")
	c := make(chan int)
	go func() {
		c <- 1
		close(c) // this is very important, everything breaks without it, still don't know why
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}

// concurrent design pattern -- fan out into as many go routines as possible,
// then fan those results back in when we get result, fan that into another channel which is just
// results
func fanIn() {
	fmt.Println("------using Fan In concurrency design------")
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send3(even, odd)           // send out work which will be written to fanin
	go receive3(even, odd, fanin) // fanin is channel of results

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func send3(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

// fan from multiple channels (even, odd) to one channel (fanin)
func receive3(e, o <-chan int, f chan<- int) {
	var wg sync.WaitGroup // initialize waitgroup, which (ensures everything gets run)

	// (add to go routines to the waitgroup)
	wg.Add(2)

	// first go routine
	go func() {
		// write every value in the even channel to the fanin channel
		for v := range e {
			f <- v
		}
		// when that's done, close this go routine
		wg.Done()
	}()
	// second go routine
	go func() {
		// write every value in the odd channel to the fanin channel
		for v := range o {
			f <- v
		}
		// when that's done, close this go routine
		wg.Done()
	}()
	// wait until all (2) go routines have finished before exiting the function
	wg.Wait()
	// close fanin channel
	close(f)
}

// taking repeat work,
// serial is one at a TIME
func fanOut() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
		fmt.Println("work # ", i)
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	// const goroutines = 10 // throttling our number of routines
	// wg.Add(goroutines)    // throttling our number of routines
	for v := range c1 {
		wg.Add(1) // if we're throttling our number of go routines,
		// then remove this bc it adds a new routine with every channel value
		// closure!!
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

// advanced topic...
func contextOverview() {
	fmt.Println("------understanding context------")
	ctx := context.Background()

	fmt.Println("context:\t", ctx)         // context.Background
	fmt.Printf("context type:\t%T\n", ctx) // *context.emptyCtx (pointer to empty context)
	fmt.Println("context err:\t", ctx.Err())

	fmt.Println("--------")
	ctx, cancel := context.WithCancel(ctx) // pass in a parent context

	fmt.Println("context:\t", ctx)         // context.Background.WithCancel
	fmt.Printf("context type:\t%T\n", ctx) // *context.cancelCtx (pointer to empty context)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Println("cancel:\t\t", cancel)       // memory address ? hex ? ' 0x106def0
	fmt.Printf("cancel type:\t%T\n", cancel) // context.CancelFunc (pointer to empty context)

	fmt.Println("--------")
	cancel()

	fmt.Println("context:\t", ctx)           // context.Background.WithCancel
	fmt.Printf("context type:\t%T\n", ctx)   // *context.cancelCtx (pointer to empty context)
	fmt.Println("context err:\t", ctx.Err()) // context cancelled !!! (change here, first non-nil)
	fmt.Println("cancel:\t\t", cancel)       // same memory address / hex
	fmt.Printf("cancel type:\t%T\n", cancel) // context.CancelFunc (pointer to empty context)
}
