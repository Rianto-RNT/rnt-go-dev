//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import (
	"fmt"
)

type Product struct {
	price int
	title string
}

func checkout(itemLists [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(itemLists); i++ {
		item := itemLists[i]
		cost += item.price

		if item.title != "" {
			totalItems += 1
		}
	}

	fmt.Println("Last item on the item list", itemLists[totalItems-1])
	fmt.Println("Total items:", totalItems)
	fmt.Println("Total cost:", cost)
}

func main() {
	cartItems := [4]Product{
		{1, "Keyboard"},
		{2, "Monitor"},
		{3, "SSD"},
	}

	checkout(cartItems)
	cartItems[3] = Product{4, "RAM"}

	checkout(cartItems)
}
