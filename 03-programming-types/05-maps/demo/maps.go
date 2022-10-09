package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	// len(shoppingList)

	shoppingList["keyboard"] = 11
	shoppingList["mouse"] = 1
	shoppingList["monitor"] += 1

	shoppingList["keyboard"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "mouse")
	fmt.Println("Mouse deleted, new list", shoppingList)

	fmt.Println("need", shoppingList["keyboard"], "keyboard")

	hardisk, found := shoppingList["hardisk"]
	fmt.Println("Need hardisk?")
	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("yup", hardisk, "items")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "items on the shopping list")
}
