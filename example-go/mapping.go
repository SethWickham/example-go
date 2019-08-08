package main

import (
	"fmt"
)

func mapping() {

	fmt.Println("mapping file Print START")
	{
		//creating a MAP with a key of TYPE STRING, adding in another record, deleting a record, then printing out all the remaining VALUES with their INDEX postion

		m := map[string][]string{

			"Wickham_Seth":     []string{"boogie bording", "walking on the beach"},
			"Wickham_Breanna":  []string{"wake bording", "makeup", "walking with Seth"},
			"Wickhams_NYA@XYA": []string{"playing with ponies", "Loving", "being LOVED"},
		}
		// this is the other record we are adding in
		m["Cousin_Ian"] = []string{"Surfing"}
		fmt.Println(m)
		// deleting our record
		delete(m, "Wickham_Seth")

		for k, v := range m {
			fmt.Println("This is the record for:", k)
			for i, v2 := range v {
				fmt.Println("\t", i, v2)
			}
		}

	}

	fmt.Println("mapping file Print END")
}
