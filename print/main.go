package main

import "fmt"

func main() {
	age := 25
	name := "ravi"
	height := 173.3                                          //float
	fmt.Println("age:", age, "height", height, "name", name) //It automatically adds spaces between operands and a newline character at the end.

	fmt.Println("helloworld") //The fmt.Printf function in Go formats data according to a format specifier and writes it to standard output.

	fmt.Printf("age is %d\n", age)         //%d is for the int
	fmt.Printf("name is %s", name)         //%s if for the string
	fmt.Printf("height is %.4f\n", height) //%.4f is forr the float
	fmt.Printf("tyoe of age is %T\n", age)
	//printing in on go
	fmt.Printf("Name: %s, age: %d, height:%.4f\n", name, age, height)
}
