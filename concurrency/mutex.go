package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int
var wgR sync.WaitGroup
var mu sync.Mutex

func raceCheck() {
	wgR.Add(2)
	go incrementor("foo race")
	go incrementor("bar race")
	wgR.Wait()
	fmt.Println("Final counter:", counter)
}

func incrementor(s string) {
	mu.Lock()
	for i := 0; i < 20; i++ {
		counter++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	}
	fmt.Println(s, " - ", counter)
	mu.Unlock()
	wgR.Done()
}
