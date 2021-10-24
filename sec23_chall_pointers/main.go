package main

import "fmt"

func main() {
	// ex 1
	x := 10.10
	// print address of x
	fmt.Printf("Address of x: %v\n", &x)
	// declare pointer to x
	ptr := &x
	// could also do
	//var ptr1 *float64
	//ptr1 = &x
	// type and value of pointer
	fmt.Printf("Type: %T\nAddress stored in pointer: %v\n", ptr, ptr)
	// address of the pointer and value it points to
	fmt.Printf("Address of pointer: %v\nAddress pointer points to: %v\nValue pointed to: %v\n",&ptr, ptr, *ptr)

	// ex 2
	x2, y2 := 10, 2
	ptrx, ptry := &x2, &y2
	// divide x2 and y2 through ptrs and store in z
	z := *ptrx / *ptry
	fmt.Printf("Value of Z - x2/y2: %v\n", z)

	// ex 3
	x3, y3 := 5.5, 8.8
	swp := mySwap
	swp(&x3, &y3)
	fmt.Printf("x3: %v\ny3: %v",x3, y3)
}

func mySwap(x *float64, y *float64) {
	store := *x
	*x = *y
	*y = store
}


