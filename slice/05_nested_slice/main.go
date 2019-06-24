package main

import "fmt"

func main() {

	//Example 1
	nested := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		out := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			out = append(out, j)
		}
		nested = append(nested, out)
	}
	fmt.Println(nested)

	//Example 2:

	appleLaptops := []string{"MacbookPro", "MacbookAir"}
	hpLaptops := []string{"hp650", "hpEliteBook"}
	laptops := [][]string{appleLaptops, hpLaptops}
	for i, v := range laptops {
		fmt.Println("Record: ", i)
		for _, name := range v {
			fmt.Printf("\t Laptop name: %v \n", name)
		}
	}

	// Example 3: Two Dimensional:
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two Dimensional: ", twoD)

}
