package main

import "fmt"

func main() {

	//learning slices

	//slices

	numbers1 := []int{1, 2, 3, 45}

	//access numbers
	fmt.Println("slice :", numbers1)

	//length of the slice
	fmt.Println("length of the slice:", len(numbers1)) //as per the numbers filled

	//append operation possible
	numbers1 = append(numbers1, 3, 6, 5, 2, 7, 89)
	fmt.Println("slice :", numbers1)

	fmt.Println("length of the slice:", len(numbers1)) //as per the numbers filled
	fmt.Printf("numbers is of type %T\n", numbers1)

	//empty slices
	name := []string{}
	println("empty slice:", name)

	// 	var name1 = [2]int{}
	// 	name1[0] = 1
	// 	fmt.Println(name1)
	//
	stock := make([]int, 0) //initializing an emty slice
	//stock := make([]int,3,5)length and the capacity of the slice

	fmt.Println("slice :", stock)
	fmt.Println("length:", len(stock))
	fmt.Println("capacity:", cap(stock))

	cash := make([]int, 3, 5)
	cash[1] = 3
	cash[2] = 2
	fmt.Println("slice 1:", cash)
	fmt.Println(append(cash, 3)) //doesnt make permanat changes
	fmt.Println(cash)
	cash = append(cash, 23)

	fmt.Println(cash)
	fmt.Println(len(cash))
	fmt.Println(cap(cash))

	//what happens if length exceeds caapcity here
	cash[0] = 4
	//cash[8] = 3//error since excceds capacity
	cash = append(cash, 44)
	cash = append(cash, 98)
	cash = append(cash, 77)
	fmt.Println(cash)
	fmt.Println(len(cash)) //according tho the numbers filled
	fmt.Println(cap(cash)) //as soon as length excced the limit the capacity doubles the older one each time
}
