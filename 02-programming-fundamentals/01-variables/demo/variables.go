package main

import "fmt"

func main() {
	var myName = "rian"

	fmt.Println("my name is: ", myName)

	var name string = "ryan"
	fmt.Println("name = ", name)

	userName := "admin"
	fmt.Println("username: ", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 54
	fmt.Println("part 1 is", part1, "other is", other)

	part2, other := 5, 21
	fmt.Println("part 2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum =", sum)

	var (
		lessonName = "Variable"
		lessonType = "Demo"
	)

	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "Hello", "RNT", "!"
	fmt.Println(word1, word2)
}
