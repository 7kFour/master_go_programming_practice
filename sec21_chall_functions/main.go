package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// ex 1
	tc := cube(14)
	fmt.Printf("The cube is: %v\n", tc)
	// ex 2
	tf, ts := f1(5)
	fmt.Printf("The factorial is: %v\nThe sum of integers: %v\n", tf, ts)
	// ex 3
	tv := myFunc("9")
	fmt.Printf("The value is: %v\n", tv)
	// ex 4
	theSum := sum(4, 8, 12, 16)
	fmt.Printf("The sum of integers: %v\n", theSum)
	// ex 5
	nSum := nakedSum(22, 34, 63, 9, 11)
	fmt.Printf("The sum (naked return) is: %v\n", nSum)
	// ex 6
	animals := []string{"shark", "spider", "dog", "raven", "scorpion"}
	result1 := searchItem(animals, "Shark")
	fmt.Printf("Is in slice: %v\n", result1)
	result2 := searchItem(animals, "wombat")
	fmt.Printf("Is in slice: %v\n", result2)
	// ex 7
	result3 := searchItemCase(animals, "Raven")
	fmt.Printf("Is in slice: %v\n", result3)
	result4 := searchItemCase(animals, "dunkey")
	fmt.Printf("Is in slice: %v\n", result4)
	// ex 8
	defer printIt("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
	// ex 9
	myInt := intFunc
	fmt.Printf("%T\n", myInt)
	myInt(33)
}

// ex 1
func cube(n float64) float64 {
	return n*n*n
}

// ex 2
func f1(n uint) (int, int) {
	fact := 1
	sum := 0
	if n < 0 {
		fmt.Println("Factorial of negative number doesn't exist!")
	} else {
		for i := 1; i <= int(n); i++ {
			fact *= i
			sum += i
		}
	}
	return fact, sum
}

//ex 3
func myFunc(s string) int {
	a, err := strconv.Atoi(s)
	_ = err
	return a + (a*11) + (a*111)
}

// ex 4
func sum(n ...int) int {
	var ts int
	for _,v := range n {
		ts += v
	}
	return ts
}

// ex 5
func nakedSum(n ...int) (s int) {
	for _, v := range n {
		s += v
	}
	return
}

// ex 6 case sensitive
func searchItem(mySlice []string, myString string) bool {
	for _, v := range mySlice {
		if v == myString {
			return true
		}
	}
	return false
}

// ex 7 case insensitive
func searchItemCase(mySlice []string, myString string) bool {
	for _, v := range mySlice {
		if v == strings.ToLower(myString) {
			return true
		}
	}
	return false
}

// ex 8
func printIt(msg string) {
	fmt.Println(msg)
}

// ex 9
func intFunc(n int) {
	fmt.Println(n)
}
