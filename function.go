package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Func add", add(42, 13))
	a, b := swap("hello", "world")

	fmt.Println("Func swap", a, b)

	x, y := split(17)
	fmt.Println("Func split", x, y)
}