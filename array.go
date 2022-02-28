package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Citra"
	names[1] = "Kemurnian"
	names[2] = "Lombu"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10,
		20,
		90,
	}

	fmt.Println(values)
}