package main

import "fmt"

func main() {

	//Function scope array:

	integerArray := [5]int{10, 20, 30, 40, 50}
	stringArray := [4]string{"first", "second", "third", "fourth"}

	fmt.Println("This is the integer array: ", integerArray)
	fmt.Println("This is the string array: ", stringArray)

}
