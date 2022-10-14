package main

import (
	"fmt"
	"time"
)

func channelTestUnBuffer() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}

func channelTestBuffer() {
	c := make(chan int, 5) //adding buffer 5

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}

func testChannelClose() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}

func testChannelUniDir() <-chan int {
	o := make(chan int)
	close(o)
	return o
}

func testDeadLockWithChannel() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

func testDeadLockSol() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)

}
