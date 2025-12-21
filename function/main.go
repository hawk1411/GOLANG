package main

import "fmt"

// func add(a int, b int) (result int) {
// 	result = a + b
// 	return
// }

func add(a int, b int) int {

	return a + b //same thing
}

func main() {
	fmt.Println("we are learning in golang")

	ans := add(2, 3)
	fmt.Println("here is your add result :", ans)

}
