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

import "fmt"

func main() {
	type product struct {
		name  string
		price float64
	}
	shoppingList := [4]product{}

	shoppingList[0] = product{name: "Milk", price: 2.50}
	shoppingList[1] = product{name: "Bread", price: 1.50}
	shoppingList[2] = product{name: "Eggs", price: 3.50}

	fmt.Println("Last item on the list:", shoppingList[2].name)
	fmt.Println("Total number of items:", len(shoppingList))
	fmt.Println("Total cost of the items:", shoppingList[0].price+shoppingList[1].price+shoppingList[2].price)

	shoppingList[3] = product{name: "Butter", price: 2.00}

	fmt.Println("Last item on the list:", shoppingList[3].name)
	fmt.Println("Total number of items:", len(shoppingList))
	fmt.Println("Total cost of the items:", shoppingList[0].price+shoppingList[1].price+shoppingList[2].price+shoppingList[3].price)

}
