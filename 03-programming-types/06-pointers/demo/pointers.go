package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}

	hello := "Hello"
	rnt := "RNT!"
	fmt.Println(hello, rnt)

	replace(&hello, "Good morning,", &counter)
	fmt.Println(hello, rnt)

	phrase := []string{hello, rnt}
	fmt.Println(phrase)

	replace(&phrase[1], "Rian!", &counter)
	fmt.Println(phrase)
}
