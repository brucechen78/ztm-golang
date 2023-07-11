//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type item struct {
	name string
	tag  bool
}

// func activateTag(item *item) {
// 	item.tag = true
// }

func deactivateTag(item *item) {
	item.tag = false
}

func checkout(items []item) {
	for i := range items {
		deactivateTag(&items[i])
	}
}

func main() {
	items := []item{
		{name: "Milk", tag: true},
		{name: "Bread", tag: true},
		{name: "Eggs", tag: true},
		{name: "Butter", tag: true},
	}

	fmt.Println("Before deactivation:", items)

	deactivateTag(&items[0])

	fmt.Println("After deactivation:", items)

	checkout(items)

	fmt.Println("After checkout:", items)
}
