package main

import "fmt"

var name string = "mojo"
var name1 int

func main() {
	fmt.Println("hello ", name)

	fmt.Printf("%v\n", name1)
	fmt.Println(name1)

	v := vrap()
	fmt.Println(v())

	fmt.Printf("%T \n", q)
	fmt.Printf("%T \n", e)
	fmt.Printf("%v\n", e+10)
}

const q string = "A"
const w string = `B`
const e = 7777

func vrap() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
