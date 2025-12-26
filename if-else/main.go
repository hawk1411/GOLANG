package main

import "fmt"

func main() {

	x := 21

	if x > 5 {
		fmt.Println("greater than 5")

	} else {
		fmt.Println("x is smaller than 5")
	}

	z := 10

	if z > 10 {
		fmt.Println("z is greater than 5 but smaller than 30")
	} else if z > 5 {
		fmt.Println("z is greater than 5 but smaller than 10")
	} else {
		fmt.Println("z is smaller than 5")
	}

	k := 23
	if k > 20 && (z > 2 || z < 3) {
		fmt.Println("hurray")
	} else {
		fmt.Println("learn pragramming hello world")
	}
}
