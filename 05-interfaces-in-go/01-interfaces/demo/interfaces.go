package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook Chicken")
}

func (c Salad) PrepareDish() {
	fmt.Println("Sepat")
	fmt.Println("Add dressing")
}

func PrepareDish(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("---Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken wings:"), Salad("Sumbawanese Salad")}
	PrepareDish(dishes)
}
