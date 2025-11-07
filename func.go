package main

import "fmt"

func add(a int, b int) int{
	return a + b
}

func sub(a int, b int) int{
	return a-b
}

func mult(a int, b int) int{
	return a*b
}

func main() {
	fmt.Println("Addition: ", add(5, 3))
	fmt.Println("Subtraction: ", sub(5, 3))
	fmt.Println("Multiplication: ", mult(5, 3))
}

