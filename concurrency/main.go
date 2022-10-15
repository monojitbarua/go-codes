package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//
	// wait group
	wg.Add(2) // adding for how many go routines and Done() calls
	go foo()
	go bar()
	wg.Wait()

	//
	//race check
	raceCheck()

	//
	//channel test
	channelTestUnBuffer()
	channelTestBuffer()
	testChannelClose()
	testChannelUniDir()
	testDeadLockWithChannel()
	testDeadLockSol()
	fanInFanOut()

}

func foo() {
	i := 0
	for i = 0; i <= 20000; i++ {
	}
	fmt.Println("Foo ", i)
	wg.Done()
}
func bar() {
	i := 0
	for i = 0; i <= 15000000; i++ {
	}
	fmt.Println("Bar ", i)
	wg.Done()
}
