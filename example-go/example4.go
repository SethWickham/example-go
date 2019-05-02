package main

import (
	"fmt"
)

// Here is an example of embedding one TYPE into another TYPE in go using a STRUCT
// in our example we are saying that our disney STRUCT is going to include TYPE hollywood and all of its FIELDS.
// then we are creating a SHORT DECLARATION called dm (of TYPE disney which also now includes TYPE hollywood)
//then we are printing out the FIELD VALUES using dot notation
type hollywood struct {
	nameOfProducer string
}

type disney struct {
	hollywood
	goofyMovie bool
}

func example4() {

	dM := disney{
		hollywood: hollywood{
			nameOfProducer: "Dan Rounds",
		},
		goofyMovie: true,
	}
	//we use dot notation to call in our FIELD VALUES, however in go our inner FIELD of nameOfProducer in our embedded TYPE hollywood
	//was promoted (meaning that all of our FIELDS are equally easy to access without excessive navigation using dot notation)
	//therefore we don't need to specify the hollywood STRUCT when trying to access the FIELD VALUE (nameOfProducer).
	fmt.Println("Name of Producer:", dM.nameOfProducer)
	fmt.Println("was The Goofy Movie made by Disney:", dM.goofyMovie)
}
