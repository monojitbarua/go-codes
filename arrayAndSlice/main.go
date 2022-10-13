package main

import (
	"fmt"
)

func main() {
	x := [10]int{10, 55, 44}
	y := []int{2000, 44444, 77}

	x[5] = 7
	y[2] = 58
	fmt.Println(x)
	fmt.Println(y)

	//
	// some array example:
	//adding int's corresponding string val in array
	var letters [58]string
	for i := 65; i < 123; i++ {
		letters[i-65] = string(i)
	}
	fmt.Println(letters)

	fmt.Println(letters[5:9])

	//
	//slice
	s := make([]int, 0, 0)
	fmt.Println("slice:", s)
	s = append(s, 5)
	s = append(s, 7)
	s = append(s, 6)
	s = append(s, 3)
	fmt.Println("slice:", s)

	//
	//delete from slice
	s = append(s[:2], s[3:]...)
	fmt.Println("slice after delete:", s)

	anotherSlice := []int{7, 2, 9, 3}
	anotherSlice1 := []int{}
	anotherSlice2 := make([]int, 40, 100)
	fmt.Println("slice:", anotherSlice)
	fmt.Println("slice1:", anotherSlice1)
	fmt.Println("slice2:", anotherSlice2)

	//
	//2D slice
	s2D := make([][]int, 4)
	s2D = append(s2D, s)
	s2D = append(s2D, anotherSlice)
	s2D = append(s2D, anotherSlice1)
	fmt.Println("2D slice: ", s2D)

	//
	//increament slice
	fmt.Println("increment slice before:", anotherSlice[1])
	anotherSlice[1]++
	fmt.Println("increment slice after:", anotherSlice[1])

}
