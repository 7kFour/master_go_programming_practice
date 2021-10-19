package main

import "fmt"

func main() {
	// ex 1
	//var m1 map[string]int // nil value map
	//_ = m1
	//
	//m2 := map[int]string{3:"Truck", 1:"Cat"} // map literal
	//_ = m2
	//
	//fmt.Printf("exiting key value: %v\n non existing key value: %v", m2[3], m1["Tractor"])

	// ex 2
	//var m1 map[int]bool
	////m1[5] = true // can't assign value to unitialized map
	//_ = m1
	//
	//m2 := map[int]int{3: 10, 4: 40}
	//m3 := map[int]int{3: 10, 4: 40}
	//
	////fmt.Println(m2 == m3) // can't compare maps with == can only compare them to nil
	//s1 := fmt.Sprintf("%d", m2)
	//s2 := fmt.Sprintf("%d", m3)
	//
	//if s1 == s2 {
	//	fmt.Println("Maps are equal")
	//} else {
	//	fmt.Println("Maps are not equal")
	//}

	// ex 3
	//m := map[int]bool{"1": true, 2: false, 3: false} // error: keys must all be the same type
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 1) // deleting 1 k:v from m

	// iterating over map to print out all k:v - this isn't super fast because
	// maps are hash tables - meant for fast lookup not fast iteration
	for k, v := range m {
		fmt.Printf("Key: %v --- Value: %v\n", k, v)
	}

}
