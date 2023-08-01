package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ex01()
	// channelsEx01()
	// selectEx()
	loopingChannels()
}

func ex01() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("this happens asynchronously")
		wg.Done()
	}()
	fmt.Println("this is synchronous")
	wg.Wait()
}

func channelsEx01() {
	var wg sync.WaitGroup
	channel := make(chan string)
	wg.Add(1)

	go func() {
		channel <- "the message"
	}()

	go func() {
		fmt.Println("message received:", <-channel)
		wg.Done()
	}()
	wg.Wait()
}

func selectEx() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "message ch1"
	}()
	go func() {
		ch2 <- "message ch2"
	}()

	time.Sleep(10 * time.Millisecond)

	select {
	case msg := <-ch1:
		fmt.Println("case 1:", msg)
	case msg := <-ch2:
		fmt.Println("case 2:", msg)
	default:
		fmt.Println("no messages")
	}
}

func loopingChannels() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
