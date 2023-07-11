package main

import "fmt"

func main() {
	var myName = "Bruce"
	fmt.Println("Hello", myName)

	var name string = "Clark"
	fmt.Println("Hello", name)

	userName := "Diana"
	fmt.Println("user is", userName)

	part1, other := 1, 2
	fmt.Println("part 1 is", part1, "other is", other)

	part2, other := 5, 4
	fmt.Println("part 2 is", part2, "other is", other)

	var sum = part1 + part2
	fmt.Println("sum is", sum)

	var (
		lessenName = "Variable"
		lessenType = "Demo"
	)
	fmt.Println("lessen name is", lessenName, "lessen type is", lessenType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)

}
