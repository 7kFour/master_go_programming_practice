package main

import "fmt"

func main() {
	// ex 1
	me := person{
		name: "Anthony",
		age: 35,
		colors: []string{"blue", "gold"},
		gr: grades{grade: 93, course: "Python"},
	}

	you := person{
		name: "Ber",
		age: 32,
		colors: []string{"panch", "white"},
		gr: grades{grade: 100, course: "Web Dev"},
	}

	fmt.Printf("%v\n%v\n", me, you)

	// ex 2
	you.name = "Andrei"
	var colors []string = you.colors
	fmt.Printf("Type: %T Value: %v", colors, colors)

	colors = append(colors, "red")
	you.colors = colors

	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

	//for i := 0; i < len(me.colors); i++ {
	//	fmt.Println("Colors:", me.colors[i])
	//}
	// better version of above
	for index, color := range me.colors {
		fmt.Printf("%d -> %q\n", index, color)
	}

	me.gr.grade = 98
	me.gr.course = "Goland"
	you.name = "Amber"

	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)
}

type person struct {
	name string
	age int
	colors []string
	gr grades
}

type grades struct {
	grade int
	course string
}

