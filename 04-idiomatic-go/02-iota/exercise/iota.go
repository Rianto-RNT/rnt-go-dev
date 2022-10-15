//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

const (
	Add = iota
	Substract
	Multiply
	Divide
)

type Operation int

func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Substract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("Unhandled operation!")
}

func main() {
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2))

	substract := Operation(Substract)
	fmt.Println(substract.calculate(20, 5))

	multiply := Operation(Multiply)
	fmt.Println(multiply.calculate(4, 4))

	divide := Operation(Divide)
	fmt.Println(divide.calculate(10, 2))
}
