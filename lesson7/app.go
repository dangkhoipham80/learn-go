package main

import (
	"fmt"
)

const PI float64 = 3.141
func main() {
	// Lesson 7
	const PI float64 = 3.14159
	const (
		ACTIVE   = 1
		INACTIVE = 0
	)
	fmt.Println("\nPI:", PI)
	fmt.Println("Status Active:", ACTIVE)
	fmt.Println("Status Inactive:", INACTIVE)
}