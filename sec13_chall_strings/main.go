package main

import "fmt"

func main () {
	// ex 1
	var name = "Anthony"
	country := "United States"
	p := fmt.Println
	fmt.Printf("My Name: %v\nCountry: %v\n", name, country)
	p(`He says: "Hello"`)
	p(`C:\Users\a.txt`)

	// ex 2
	// a with caron
	r := 'ǎ'
	fmt.Printf("r is type: %T\n", r)
	s1, s2 := "ma", "m"
	p(s1 + s2 + string(r))
}