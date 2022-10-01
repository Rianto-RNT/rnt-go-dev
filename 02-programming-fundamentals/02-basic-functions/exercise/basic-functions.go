//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetings(n string) {
	fmt.Println("Good Morning,", n)
}

func sayings() string {
	return "I like go programming"
}

func addNum(a, b, c int) int {
	return a + b + c
}

func sevenTeen() int {
	return 17
}

func num2() (int, int) {
	return 17, 71
}

func main() {
	greetings("Ryan")

	fmt.Println(sayings())

	x, y := num2()
	result := addNum(sevenTeen(), x, y)
	fmt.Println(result)
}
