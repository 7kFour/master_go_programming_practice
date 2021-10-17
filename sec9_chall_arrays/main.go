package main

import "fmt"

func main() {
	// ex 1
	var cities [2]string // will just be 2 empty values
	fmt.Println(cities)

	// ex 1.2 - array literal - 2 examples of array literal initialization
	grades := [3]float64{92.6, 53.4}
	var moreGrades = [3]float64{73.2, 87.6}
	fmt.Printf("Grades: %#v\n", grades)
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("More grades: %#v\n", moreGrades)
	fmt.Printf("More grades: %v\n", moreGrades)

	// ex 1.3 - just says to declare - not initialize or use
	taskDone := [...]bool{}
	_ = taskDone // blank identifier use

	// ex 1.4
	for i := 0; i < len(cities); i++ {
		fmt.Println("Cities array: " + cities[i])
	}

	// ex 1.5
	for i, v := range grades {
		fmt.Println("Index:", i, "Value:", v)
	}

	// ex 2
	nums := [...]int{30, -1, -6, 90, -6, 27}
	var count int
	for _, v := range nums {
		if v >= 0 && v % 2 != 1 {
			count++
		}
	}
	fmt.Println("Count of unsigned and even:", count)

	// ex 3
	myArray := [3]float64{1.2, 5.6}
	myArray[2] = 6

	//a := 10
	a := 10.10 // making this a float will correct the error that we were getting below
	myArray[0] = a // error here because trying to assign int to float

	// commented out because it would go out of bounds in the array
	//myArray[3] = 10.10

	fmt.Println(myArray)
}
