package main

import "fmt"

func fanInFanOut() {
	i := insertChannel()
	o := channelOut(i)
	for v := range o {
		fmt.Println(v)
	}

}

func insertChannel() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 1; j < 11; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func channelOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
