package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.
	switch {
	case runtime.GOOS == "darwin":
		fmt.Println("OS X.")
	case runtime.GOOS == "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", runtime.GOOS)
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool.")
		case int:
			fmt.Println("I'm an int.")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
}