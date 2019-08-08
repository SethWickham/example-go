package main

import "fmt"

// this file contains some other examples of go looping statements and mapping
func loops() {
	fmt.Println("Case Statements Print START:")

	//for loop: First we declare our variable and initialize
	//our loop then we have the condition that will repeat
	//until its no longer evaluated as true.
	//INIT, CONDITION, POST
	zToTen := 0
	for zToTen < 10 {
		fmt.Println(zToTen)
		zToTen++
	}
	fmt.Println(zToTen)

	// if both statements are true it will run Println
	if 2 != 3 && 2 != 4 {
		fmt.Println("two is not equal to three AND 2 is not equal to 4")
	}
	//if either statement is true then it will run Println
	if 2 == 2 || 3 == 3 {
		fmt.Println("two is equal to two, OR three is equal to three")
	}
	// statements are automatically evaluated as true unless otherwise specified
	if true {
		fmt.Println("Print me I'm true")
	}
	//this statement will not Print because its evaluated as false
	if false {
		fmt.Println("nope do not Print")
	}

	fmt.Println("Case Statements Print END:")
}
