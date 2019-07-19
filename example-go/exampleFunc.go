package main

import "fmt"

//this file contains demonstrations of functions in go

func exampleFunc() {
	//this is a what a functions looks like in go
	// func (r reciever) identifier(parameters)(return(s)) { code }
	demoInt(1, 2, 3, 4, 5)
	demoStr("Johnny Appleseed")
}

//using the ... operator
//we are able to RECIEVE an unlimited number of INT into our FUNCTION (r reciever).
//the scope of s in our func is only within the curly braces
//therefore we can use s in both of our functions (r reciever)
func demoInt(s ...int) {
	fmt.Println("example3 file print int", s)
}

func demoStr(s string) {
	fmt.Println("example3 file print string", s)

}
