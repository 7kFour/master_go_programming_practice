package main

import (
	"fmt"
)

// ex 1
type money float64

// ex 3
type book struct {
	title string
	price float64
}

func main() {
	// ex 1
	var theM money = 123.4567
	fmt.Printf("Amount: %v\n", theM)
	fmt.Printf("The mf amount: ")
	theM.print()
	fmt.Println() // adding newline

	// ex 2
	var salary money = 1.5 * 5.503
	fmt.Println("The salary:", salary)
	salary.print()
	fmt.Println() // adding newline
	s := salary.printStr()
	fmt.Println(s)

	// ex 3
	bestBook := book{title: "Dune", price: 9.9}
	vat := bestBook.vat()
	fmt.Printf("Vat: %v\n", vat)
	bestBook.discount()
	fmt.Printf("%#v\n", bestBook)

	// ex 4
	bestBook.changePrice(11.99)
	fmt.Printf("Book's price:%.2f\n", bestBook.price)

}
// ex 1
// method name is print
func (n money) print() {
	fmt.Printf("%.2f", n)
}
// ex 2
func (n money) printStr() string {
	return fmt.Sprintf("%.2f",n)
}
// ex 3
func (b book) vat() float64 {
	return b.price * 0.09
}
func (b *book) discount() {
	// use a pointer receiver to change the struct value
	(*b).price = b.price * 0.9
	// equivalent to:
	// b.price = b.price * 0.9
}
// ex 4
func (b *book) changePrice(p float64) {
	(*b).price = p
}
