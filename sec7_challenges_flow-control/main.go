package main

import (
	"fmt"
	"time"
)

func main() {
	// ex 1
	for i := 0; i < 50; i++ {
		if i % 7 == 0 {
			fmt.Println(i);
		}
	}

	// ex 2
	for i := 0; i < 50; i++ {
		if i % 7 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// ex3
	count := 0
	for i := 0; i < 50; i++ {
		if i % 7 == 0 {
			fmt.Println(i)
			count++
		}
		if count == 3 {
			break
		}
	}

	// ex 4
	for i := 0; i < 501; i++ {
		if i % 7 == 0 && i & 5 == 0 {
			fmt.Println(i)
		}
	}

	// ex 5
	// getting the current year
	//ty := time.Now().Year()
	//for i := 1985; i < ty; {
	//	fmt.Println(i)
	//	i++
	//}
	// refactor above
	by, ty := 1985, time.Now().Year()
	for i := by; i < ty; {
		fmt.Printf("%v\n", i)
		i++
	}

	// ex 6
	age := -9
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age.")
	case age < 18:
		fmt.Println("You are a minor!")
	case age == 18:
		fmt.Println("Congrats, you're an adult!")
	default:
		fmt.Println("You're an adult.")
	}
}
