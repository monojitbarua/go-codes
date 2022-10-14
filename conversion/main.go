package main

import (
	"fmt"
	"strconv"
)

func main() {

	//int to string
	s := strconv.Itoa(-52)
	fmt.Printf("%T: ", s)
	fmt.Println(s)

	//string to float
	f, _ := strconv.ParseFloat("8.88", 64)
	fmt.Printf("%T: ", f)
	fmt.Println(f)

	//string to bool
	b, _ := strconv.ParseBool("false")
	fmt.Printf("%T: ", b)
	fmt.Println(b)

}
