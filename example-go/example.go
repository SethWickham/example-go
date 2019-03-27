package main

import (
	"fmt"
)

func main() {
	//Go Pillars Efficient Execution, Efficient compilation, and ease of programming
	// a SLICE allows you to group together VALUES of the same TYPE

	// COMPOSITE LITERAL --  []type{value}
	cl := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println("COMPOSITE LITERAL PRINT:", cl)

	// an example of a Multi Dimensional Slice
	iceCream := []string{"Chocolate", "Vanilla", "Strawberry"}
	fmt.Println("first string slice:", iceCream)

	// using the built in append feature to add another VALUE to my SLICE
	iceCream = append(iceCream, "Hazelnut")
	fmt.Println("append Print:", iceCream)

	numberString := []string{"32", "64", "8"}
	fmt.Println("second string slice:", numberString)
	// SLICE of a SLICE
	xi := [][]string{iceCream, numberString}
	fmt.Println("slice of a slice:", xi)
	// Mapping
	m := map[string]int{
		"Seth":    32,
		"ShirtId": 4567,
	}
	fmt.Println(m)
	fmt.Println(m["Seth"])
	fmt.Println(m["Jeans"])

	//The comma ok idiom
	if v, ok := m["Jeans"]; ok {
		fmt.Println("THIS IS THE IF PRINT FOR A VALUE THAT DOESN't EXIST, Value:", v)
	} else if v, ok := m["Seth"]; ok {
		fmt.Println("THIS IS THE IF PRINT FOR A VALUE THAT EXISTS, Value:", v)

		// Pulling in a method on a seperate file but in the same package looks like this:
		other()

		//My name in binary using conversion to convert a STRING into SLICE OF BYTE
		x := "Seth Wickham"
		nx := []byte(x)
		fmt.Println("My name normally:", x)

		fmt.Printf("If a computer called out my name with no translator I would hear this:%b\n", nx)

	}
}
