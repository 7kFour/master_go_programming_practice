package main

import (
	"fmt"
)

func main() {

	// ex 1
	//fmt.Println("Exercise 1:")
	//slc := []string{"Red", "Green", "Blue"}
	//for i, v := range slc {
	//	fmt.Printf("Index: %v\tValue: %v\n", i, v)
	//}

	// ex 2
	//fmt.Println("Exercise 2:")
	//mySlice := []float64{1.2, 5.6}
	//mySlice[1] = 6
	//
	//a := 10
	//mySlice[0] = float64(a)
	//
	//mySlice[1] = 10.10
	//
	//mySlice = append(mySlice, float64(a))
	//fmt.Println(mySlice)

	// ex 3
	//fmt.Println("Exercise 3:")
	//nums := []float64{2.6, 13.3, 15.87}
	//nums = append(nums, 10.1)
	//nums = append(nums, 4.1, 5.5, 6.6)
	//fmt.Println(nums)
	//
	//n := []float64{22.3, 67.5}
	//nums = append(nums, n...)
	//fmt.Println(nums)

	// ex 4
	//fmt.Println("Exercise 4:")
	//// printing error if not run with arguments
	//if len(os.Args) < 2 {
	//	log.Fatal("Please run the program with arguments (2-10 numbers)!")
	//}
	//
	//// creating new slice of arguments
	//args := os.Args[1:]
	//
	//// declare and initialize sum and product of float64 args
	//sum, product := 0., 1.
	//
	//if len(args) < 2 || len(args) > 10 {
	//	fmt.Println("Please enter between 2 and 10 numbers!")
	//} else {
	//	for _, v := range args {
	//		num, err := strconv.ParseFloat(v, 64)
	//		if err != nil {
	//			// if string can't be converted to float64 continue to next arg
	//			continue
	//		}
	//		sum += num
	//		product *= num
	//	}
	//}
	//fmt.Printf("Sum: %v, Product: %v\n", sum, product)

	// ex 5
	//fmt.Println("Exercise 5:")
	//nums5 := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	//sum5 := 0
	//// this ignores the first two and last two elements
	//for _, v := range nums5[2 : len(nums5)-2] {
	//	fmt.Println(v)
	//	sum5 += v
	//}
	//fmt.Printf("Sum: %v\n", sum5)

	// ex 6
	//fmt.Println("Exercise 6:")
	//friends := []string{"Marry", "John", "Paul", "Diana"}
	//// creating new slice (type, length)
	//fscp := make([]string, len(friends))
	//// copying from friends to fscp (dst,src)
	//copy(fscp, friends)
	//// proving that changing one slice doesn't affect the other
	//fscp[3] = "Gayorg"
	//fmt.Printf("friends: %v\nfscp: %v", friends, fscp)

	// ex 7
	//fmt.Println("Exercise 7:")
	//friends := []string{"Marry", "John", "Paul", "Diana"}
	//fsapd := []string{}
	//// creating copy of friends using append
	//fsapd = append(fsapd, friends...)
	//// proving that changing one doesnt affect the other
	//friends[0] = "Leron"
	//fmt.Printf("friends: %v\nfsapd: %v", friends, fsapd)

	// ex 8
	fmt.Println("Exercise 8:")
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	// appending the first 3 and last 3 elements of years to newYears
	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Printf("%v", newYears)
}
