package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// fmt.Println("hey, what's your name?by Scan")
	// var name string

	// fmt.Scan(&name) //gave the memory refernce
	// fmt.Printf("hello, Mr. %s\n", name)

	fmt.Println("hello boi by bufio.reader")

	reader := bufio.NewReader(os.Stdin) //bufio.NewReader(os.stdin)
	name1, _ := reader.ReadString('\n') //demlimiter read till the new line
	fmt.Println("hello, Mr.", name1)

}
