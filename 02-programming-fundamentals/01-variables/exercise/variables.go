//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favoriteColor = "Black"
	fmt.Println("My favorite color:", favoriteColor)

	year, ageYear := 1993, 29
	fmt.Println("My birth year:", year, "My age:", ageYear)

	var (
		firstName = "Ryan"
		lastName  = "Morrison"
	)
	fmt.Println("My name is:", firstName, lastName)

	var ageDay int
	ageDay = 356 * ageYear
	fmt.Println("i'm", ageDay, "years old (<<< in days)")
}
