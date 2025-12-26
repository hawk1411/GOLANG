package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	age       int
}

//structure under structure

type address struct {
	//storing variables
	city string
	town string
}

type contact_details struct {

	//variables
	Phone_no int
	email    string
}

// combined structure

type employee struct {
	person_details Person
	person_address address
	person_contact contact_details
}

func main() {

	//object decleration

	var Prince Person
	fmt.Println("prince person :", Prince) //" " " " 0//____0"
	Prince.Firstname = "prince"
	Prince.Lastname = "aggarwal"
	Prince.age = 24
	fmt.Println("prince person :", Prince)

	//2nd method
	person1 := Person{
		Firstname: "prince1",
		Lastname:  "aggarwal",
		age:       23,
	}

	fmt.Println(person1)

	//3rd method

	var person3 = new(Person) //through memory refernce
	person3.Firstname = "simran"
	person3.Lastname = "agarwal"
	person3.age = 26

	fmt.Println("person 3 :", person3)
	fmt.Println("person 3 :", person3.Firstname)

	//use of the combined structure
	var Employee employee
	Employee.person_address = address{
		city: "byrnihat",
		town: "lumboi",
	}
	Employee.person_contact = contact_details{
		Phone_no: 1234345455,
		email:    "ivar",
	}

	Employee.person_details = Person{
		Firstname: "ravi",
		Lastname:  "mishra",
		age:       21,
	}

	fmt.Println("structure under structure:", Employee)
	//changing any thing
	Employee.person_details.Firstname = "RAVI"
	fmt.Println("chnged the first name :", Employee.person_details.Firstname)

}
