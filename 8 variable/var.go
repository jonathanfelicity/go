package main

import "fmt"

func main() {
	// explicite -> with type
	var email string = "jonathanfelicity@fake.com"

	// implicite -> without type
	var full_name = "Jonahtan Felicity"
	age := 45



	// we can also do it this way 

	// var( 
	// 		name string
	// 		age int8
	// 		isGood bool
	// )
	// var(
	// 	naMe "Jonathan Felicity" 
	// 	45 
	// 	true)

	fmt.Printf("%s %s %d \n", email, full_name, age)
}
