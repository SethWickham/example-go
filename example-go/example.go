package main

import (
	"fmt"
)

func main() {
	//Go Pillars Efficient Compilation, Efficient Execution, and ease of programming
	// a SLICE allows you to group together VALUES of the same TYPE

	// COMPOSITE LITERAL --  []type{value}
	cl := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println("COMPOSITE LITERAL print:", cl)

	//Locating and Printing VALUES found  within a SLICE by finding their position on a zero based index
	findx := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("First Index find Print: ", findx[0:5])
	fmt.Println("Second Index find Print: ", findx[0:5][7:9])

	// an example of a Multi Dimensional Slice
	iceCream := []string{"Chocolate", "Vanilla", "Strawberry"}
	fmt.Println("first string slice:", iceCream)

	// using the built in append feature to add another VALUE to my SLICE
	iceCream = append(iceCream, "Hazelnut")
	fmt.Println("append Print:", iceCream)

	numberString := []string{"32", "64", "8"}
	fmt.Println("second STRING SLICE:", numberString)
	// SLICE of a SLICE
	ss := [][]string{iceCream, numberString}
	fmt.Println("SLICE of a SLICE:", ss)
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

		//Deleting ShirtId
		delete(m, "ShirtId")
		fmt.Println("Print after delete:", m)
		//checking to make sure ShirtId was deleted using the if ok idiom
		if v, ok := m["ShirtId"]; ok {
			fmt.Println("delete didn't work", v)
		}

		// Pulling in a method on a seperate file but in the same package looks like this:
		other()

		//My name in binary and ascii using conversion to convert a STRING into SLICE OF BYTE
		x := "Seth Wickham"
		nx := []byte(x)
		fmt.Println("My name normally:", x)

		//my name in Binary
		fmt.Printf("If a computer called out my name with no translator I would hear this:%b\n", nx)

		//my name in ascii characters
		fmt.Printf("my name in ascii characters: %#U", nx)

	}
}
