package main

import "fmt"

// ex 2
type empty interface {}

// ex 1
type vehicle interface {
	License() string
	Name() string
}
// ex 1
type car struct {
	licenseNo string
	brand string
}
// ex 3
type cube struct {
	edge float64
}
// ex 1
func (c car) License() string {
	return c.licenseNo
}
func (c car) Name() string {
	return c.brand
}
// ex 3
func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	// ex 1
	var vh vehicle = car{brand: "Ford Mustang", licenseNo: "POW100ZZ"}
	fmt.Println(vh.License())
	fmt.Println(vh.Name())

	// ex 2
	var empty interface{}
	empty = 10
	fmt.Printf("empty is type - %T\n", empty)
	empty = 23.3
	fmt.Printf("empty is now type - %T\n", empty)
	empty = []int{1,2,3}
	fmt.Printf("empty is now type - %T\n", empty)
	empty = append(empty.([]int), 10)
	fmt.Printf("%v\n", empty)

	// ex 3
	var x interface{}
	x = cube{edge: 5}
	// type assertion
	v := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", v)
}
