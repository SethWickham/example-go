package main

import (
	"fmt"
)

func main() {
	//Go Pillars Efficient Compilation, Efficient Execution, and ease of programming
	//Go automatically adds the semi colons if we want to write multiple expressions on the same line then just add our own semicolon

	// COMPOSITE LITERAL --  []type{value}
	cl := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println("COMPOSITE LITERAL print:", cl)

	//Locating and Printing VALUES found  within a SLICE by finding their position on a zero based index
	//we can also use this for DELETING VALUES from our SLICE of INT
	findx := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("First Index find Print: ", findx[:5])
	fmt.Println("Second Index find Print: ", findx[:5][7:9])
	fmt.Println("Third Index find Print: ", findx[4:])
	// using the index to range through two arguments
	//we are taking the 0-3 index and the 6-end index spots and removing 45, 46, 47
	findx = append(findx[:3], findx[6:]...)
	fmt.Println("DELETE Print for findx: ", findx)

	// a SLICE allows you to group together VALUES of the same TYPE
	// an example of a Multi-Dimensional SLICE
	// iceCream SLICE
	iceCream := []string{"Chocolate", "Vanilla", "Strawberry"}
	fmt.Println("first string slice:", iceCream)

	// using the built in APPEND feature to add another VALUE to my SLICE
	iceCream = append(iceCream, "Hazelnut")
	fmt.Println("append Print:", iceCream)

	numberString := []string{"32", "64", "8"}
	fmt.Println("second STRING SLICE:", numberString)

	//using the ... feature in Go to dump in all the VALUES of our SLICE
	iceCream = append(iceCream, numberString...)
	fmt.Println("The iceCream and numberString Print:", iceCream)
	// SLICE of a SLICE using [][] shows us that we are merging together multiple SLICES, string tells us what type they are of and then we are passing in the values  which are located inside curly brackets (our two slices of string named: iceCream and numberString)
	veggieString := []string{"Carrots", "Broccoli", "Celery"}
	//This is the TYPE SLICE of a SLICE of STRING
	allStrings := [][]string{veggieString, iceCream}
	fmt.Println("SLICE of a SLICE:", allStrings)

	// Mapping
	m := map[string]int{
		"Seth":    32,
		"ShirtId": 4567,
	}
	fmt.Println(m)
	fmt.Println(m["Seth"])
	//This will not print an error even though there is no "Jeans" Property so therefore we can check it with the if statement below
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

		//My name in binary and ascii using conversion to convert a STRING into SLICE OF BYTE
		x := "Seth Wickham"
		nx := []byte(x)
		fmt.Println("My name normally:", x)

		//my name in Binary
		fmt.Printf("If a computer called out my name with no translator I would hear this:%b\n", nx)

		//my name in ascii characters
		fmt.Printf("my name in ascii characters: %#U", nx)

	}
	// calling our methods within the same package but seperate file:
	caseStatements()
	exampleFunc()
	example4()

}
