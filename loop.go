package main

import (
	"fmt"
)

func main() {
	// for loop
	fmt.Println("for loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	fmt.Println("while loop")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	fmt.Println("infinite loop")
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}

	// loop with continue
	fmt.Println("loop with continue")
	for m := 0; m < 5; m++ {
		if m == 2 {
			continue
		}
		fmt.Println(m)
	}
}