//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func humanAge() int {
	return 18
}

// const (
// 	NewBorn  = 0
// 	Toddler  = 1
// 	Child    = 2
// 	Teenager = 3
// 	Adult    = 4
// )

func main() {
	switch h := humanAge(); {
	case h == 0:
		fmt.Println("Human Newborn")
	case h >= 1 && h <= 3:
		fmt.Println("Human Toddler")
	case h < 13:
		fmt.Println("Human Child")
	case h < 18:
		fmt.Println("Human Teenager")
	default:
		fmt.Println("Human Adult")
	}
}
