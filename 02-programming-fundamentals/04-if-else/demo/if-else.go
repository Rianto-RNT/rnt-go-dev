package main

import "fmt"

func average(a, b, c int) float32 {
	// conver the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 6, 7

	if quiz1 > quiz2 {
		fmt.Println("quiz 1 score is higer than quiz 2")

	} else if quiz1 < quiz2 {
		fmt.Println("Quiz 2 sore is higer than Quiz 1")
	} else {
		fmt.Println("Quiz 2 and Quiz 2 have the same score")
	}

	if average(quiz1, quiz2, quiz3) > 6 {
		fmt.Println("Acceptable grades")
	} else {
		fmt.Println("Instructor did a bad job!")
	}
}
