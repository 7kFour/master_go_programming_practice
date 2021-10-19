package main

import "fmt"

func main () {
	// ex 1
	//var name = "Anthony"
	//country := "United States"
	//p := fmt.Println
	//fmt.Printf("My Name: %v\nCountry: %v\n", name, country)
	//p(`He says: "Hello"`)
	//p(`C:\Users\a.txt`)

	// ex 2
	// a with caron
	//r := 'ǎ'
	//fmt.Printf("r is type: %T\n", r)
	//s1, s2 := "ma", "m"
	//p(s1 + s2 + string(r))

	// ex 3
	//s1 := "țară means country in Romanian"
	//fmt.Printf("Bytes in string: ")
	//for i := 0; i < len(s1); i++ {
	//	fmt.Printf("%v", s1[i])
	//}
	//fmt.Println()
	//// iterating over the string and print out rune by rune
	//// and the byte index where the rune starts in the string
	//for i, r := range s1 {
	//	fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	//}
	//fmt.Println()

	// ex 4
	//s1 := "Go is cool!"
	//x := s1[0]
	//fmt.Println(x)
	//
	//// strings are immutable - cant do this
	////s1[0] = 'x'
	//
	//// printing the number of runes in the string
	//fmt.Println(utf8.RuneCountInString(s1))

	// ex 5
	//s := "你好 Go!"
	//// converting string to byte slice
	//b := []byte(s)
	//
	//// printing out the byte slice
	//fmt.Printf("%#v\n", b)
	//
	//// iterating over the byte slice and printing out each index and byte in the byte slice
	//for i, v := range b {
	//	fmt.Printf("%d -> %d\n", i, v)
	//}

	// ex 6
	s := "你好 Go!"

	// converting string to rune slice
	r := []rune(s)

	// printing out the rune slice
	fmt.Printf("%#v\n", r)

	// iterating over the rune slice and printing out each index and rune in the rune slice
	for i, v := range r {
		fmt.Printf("%d -> %c\n", i, v)
	}
}