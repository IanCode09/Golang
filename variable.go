package main

import "fmt"

func main() {
	var name string

	name = "Ian Lombu"
	fmt.Println(name)

	var age = 26
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Ian"
		lastName = "Lombu"
	)
	fmt.Println(firstName + " " + lastName)
}