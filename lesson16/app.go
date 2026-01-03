package main

import "fmt"

func operationPrint(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("Sum:", sum)
}

func operation(num1 int, num2 int) int {
	return num1 + num2
}
func main() {
	operationPrint(3, 5)

	fmt.Print(operation(2, 4))
}