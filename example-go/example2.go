package main

import "fmt"

// this file contains some other examples of go case statements and mapping
func example2() {
	fmt.Println("I'm the example2 file print")

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

	//case statements
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
	//this statement will not Print becuase its evaluated as false
	if false {
		fmt.Println("nope do not Print")
	}
	{
		//creating a MAP with a key of TYPE STRING, adding in another record, deleting a record, then printing out all the remaining VALUES with their INDEX postion

		m := map[string][]string{

			"Wickham_Seth":     []string{"boogie bording", "walking on the beach"},
			"Wickham_Breanna":  []string{"wake bording", "makeup", "walking with Seth"},
			"Wickhams_NYA@XYA": []string{"playing with ponies", "Loving", "being LOVED"},
		}

		m["Wickham_ian"] = []string{"electric boogaloo"}
		fmt.Println(m)

		delete(m, "Wickham_ian")

		for k, v := range m {
			fmt.Println("This is the record for:", k)
			for i, v2 := range v {
				fmt.Println("\t", i, v2)
			}
		}

	}
}
