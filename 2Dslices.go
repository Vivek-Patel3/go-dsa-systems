package main

import "fmt"

func twoD() {
	var rows int = 5
	var cols int = 3

	var multiD = make([][]int, rows)

	for i := range multiD {
		multiD[i] = make([]int, cols)
	}

	fmt.Println(multiD)
}
