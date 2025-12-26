package main

import "fmt"

func main(){

	//name <--> grade//unordered

	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 100
	studentGrades["alice"] = 90
	studentGrades["bob"] = 85

	fmt.Println("marks of the bob:", studentGrades["bob"])

	fmt.Println("marks of the sassy:", studentGrades["sassy"])

	// ucan directly update using the keys'

	studentGrades["bob"] = 100
	fmt.Println("marks of the bob:", studentGrades["bob"]) //updated to hundred

	delete(studentGrades, "bob")
	fmt.Println("marks of the bob:", studentGrades["bob"]) //updated to hundred

	//checking if a key exists
	grades, exiss := studentGrades["sassy"]
	fmt.Println("her marks", grades)
	fmt.Println("does she exists", exiss)

	studentGrades["sassy"] = 200
	grades1, exiss1 := studentGrades["sassy"]
	fmt.Println("her marks", grades1)
	fmt.Println("does she exists", exiss1)


	for index, value := range studentGrades{

		fmt.Printf("key is %s and marks is %d\n", index,value)
	}
   


	person := map[string]int{

		"alice": 90,
		"bob": 85,
		"charlie": 95, 


	}


	for index,value := range person{ 

	
		fmt.Printf("thr kry is %s and the marks is %d\n", index, value)
	} 
} 
