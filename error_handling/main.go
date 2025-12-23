package main

import "fmt"

// func divide(a, b int) int {
// 	return a / b
// }

// func divide(a,b float64) float64{
// 	if b ==0{
// 		return 0
// 	}
// 	return a/b
// }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cant be zero dudu") //or a direct string  will also doo
	}
	return a / b, nil
}

func main() {

	fmt.Println("started Error handling in the functions")
	ans, _:= divide(10, 0)//ignoring the error
	fmt.Println("answer is;", ans)


	ans1, err:= divide(10,0)

	if  err != nil{
		fmt.Println("error handling")
		fmt.Println(err)
	
}
   fmt.Println(ans1)
}
