package main

import "fmt"

//creating a function to modify the value of the variable via pointer

func modifybyreference(num *int) { //a pointer with the int type will be taken
	*num = *num * 10
	//changed directly at the address

}

func main() {

	//pointer
	var num int
	num = 2

	var ptr *int
	ptr = &num //passing the address of the num to the pointer

	fmt.Println(ptr) //an address

	//other way around

	number := 87

	pointer := &number

	fmt.Println(number, pointer)

	// modifybyreference(pointer)
	modifybyreference(&number) //same thing
	fmt.Println(number)        //870*10=87

}
