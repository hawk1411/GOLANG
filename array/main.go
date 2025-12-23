package main

import "fmt"

func main() {

	//creating an array
	//size is fixed
	var name [5]string //size is 5,2 empt strings in the end
	name[2] = "princr"
	name[1] = "ravi"
	name[0] = "aaksah"
	fmt.Println("name is: ", name)
	fmt.Printf("viewing the empty strings with cotations %q\n", name)

	var numbers = [8]int{1, 2, 3, 4, 5} //size is 8
	fmt.Println("numbers are:", numbers)

	fmt.Println("length of the array:", len(numbers)) //as per the fixed size declared

}
