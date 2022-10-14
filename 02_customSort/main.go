package main

import (
	"fmt"
	"sort"
)

// string array sort
type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

//

// custom fields sort
type emp struct {
	name string
	sal  float64
}
type employees []emp

func (p employees) Len() int           { return len(p) }
func (p employees) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p employees) Less(i, j int) bool { return p[i].sal < p[j].sal }

func main() {

	// slice of string sort
	p := people{"YYY", "A", "AA", "XX", "X"}
	fmt.Println("Before sort: ", p)
	sort.Sort(p)
	fmt.Println("After sort: ", p)

	fmt.Println("Before reverse sort: ", p)
	sort.Sort(sort.Reverse(p))
	fmt.Println("After reverse sort: ", p)

	t := sort.StringSlice(p)
	fmt.Println("slice string: ", t)
	fmt.Printf("slice string type: %T\n", t)

	//
	//custom complex field sort
	sliceOfEmployee := employees{
		emp{name: "Monojit", sal: 30.0},
		emp{name: "Nilay", sal: 60.0},
		emp{name: "Rajeev", sal: 90.0},
		emp{name: "Prashant", sal: 40.0},
	}
	fmt.Println("Before custom sort: ", sliceOfEmployee)
	sort.Sort(sliceOfEmployee)
	fmt.Println("After custom sort: ", sliceOfEmployee)
}
