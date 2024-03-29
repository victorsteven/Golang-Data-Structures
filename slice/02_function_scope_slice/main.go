package main

import "fmt"

func main() {

	//Function scope slice:

	//EXAMPLE 1:
	integerSlice := []int{10, 20, 30, 40, 50}
	stringSlice := []string{"first", "second", "third", "fourth"}
	fmt.Println("This is the integer slice: ", integerSlice)
	fmt.Println("This is the string slice: ", stringSlice)

	// EXAMPLE 2:
	// We can also define slice using the "make" builtin function
	s := make([]int, 4)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40
	fmt.Println("Slice created with 'make': ", s)

	//a. We can add elements to this slice using "append" builtin function:
	s = append(s, 50)
	fmt.Println("Added one element to slice: ", s)

	//b. We can append more than one element to the slice:
	s = append(s, 60, 70)
	fmt.Println("Added two elements to slice: ", s)

	//c. We can remove from that slice:
	//Say we want to remove "30" which is the third element at index "2"
	s = append(s[:2], s[2+1:]...) //... is a variadic argument in Go
	fmt.Println("Deleted one element from slice: ", s)

	//d. Replace an element with another
	// Say i want replace the third element now "40" with the fifth element "60":
	s[2] = s[len(s)-2]
	fmt.Println("Slice with element replaced: ", s)

	// Say i want replace the third element now "60" with the last element "70":
	s[2] = s[len(s)-1]
	fmt.Println("Updated Slice with element replaced: ", s)

	//Get particular elements from the slice:
	//We now have this slice: [10 20 70 50 60 70]
	//To get the 2nd to the 4th element, we do:
	s = s[1:4]
	fmt.Println("Slice with second to fourth element: ", s)

}
