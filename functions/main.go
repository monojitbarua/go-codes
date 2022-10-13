package main

import "fmt"

func main() {
	f1 := foo("Monojit", "Barua")
	f2, f3 := foo2("Monojit", "Barua")
	fmt.Println(f1)
	fmt.Println(f2, f3)

	//
	//variadics
	fmt.Println(sumVariadics(2, 3, 4, 5, 6))
	a := []int{7, 8, 9, 6, 5, 4}
	fmt.Println(sumVariadics(a...)) //calling variadics by unfurling slice

	//
	//func as var
	funcVar := func() {
		fmt.Println("Anonymous func call as var")
	}
	funcVar()
	funcVarTest(funcVar)

	//
	//recurrsion test with factorial
	fmt.Println("recurrsion test with factorial: ", recurssionFactorial(4))

	//
	//defer test: before the callee method's end execution defer method gets called.
	//Here deferFunc1() gets called before main() execution ends (so by default also after deferFunc2())
	defer deferFunc1()
	deferFunc2()

	//
	//passbyvalue proof: without pointer value update from function wont work
	x := "old value"
	passByValue1(x)
	fmt.Println("from passByValue1: ", x)
	y := "old value"
	passByValue2(&y)
	fmt.Println("from passByValue1: ", y)

	//
	x1 := true
	x2 := false
	x3 := x1 && x2
	fmt.Println(x3)
}

func foo(fname, sname string) string {
	return fmt.Sprint(fname, sname)
}
func foo2(fname, sname string) (string, string) {
	return fmt.Sprint(fname, sname), fmt.Sprint(fname, sname)
}

func sumVariadics(x ...int) int {
	sum := 0
	fmt.Printf("%T\n", x) //variadics input automcatically becames slice of input type
	for _, i := range x {
		sum += i
	}
	return sum
}

func funcVarTest(funcVar func()) {
	funcVar()

}

func recurssionFactorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * recurssionFactorial(x-1)
}

func deferFunc1() {
	fmt.Println(" defer func 1")
}
func deferFunc2() {
	fmt.Println(" defer func 2")
}

func passByValue1(x string) {
	x = "updated new value"
}

func passByValue2(x *string) {
	*x = "updated new value"
}
