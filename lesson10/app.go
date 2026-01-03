package main

import "fmt"

func main() {
	// Lesson 10: Arithmetic Operations
	s1, s2 := 15, 4

	sum := s1 + s2
	diff := s1 - s2
	prod := s1 * s2
	quot := float32(s1) / float32(s2)
	rem := s1 % s2

	fmt.Printf("Sum: %d + %d = %d\n", s1, s2, sum)
	fmt.Printf("Difference: %d - %d = %d\n", s1, s2, diff)
	fmt.Printf("Product: %d * %d = %d\n", s1, s2, prod)
	fmt.Printf("Quotient: %d / %d = %.2f\n", s1, s2, quot)
	fmt.Printf("Remainder: %d %% %d = %d\n", s1, s2, rem)

	// &, ==, !
	fmt.Printf("\nEqual: %d == %d is %t\n", s1, s2, s1 == s2)
	fmt.Printf("Not Equal: %d != %d is %t\n", s1, s2, s1 != s2)
	fmt.Printf("Bitwise AND: %d & %d = %d\n", s1, s2, s1&s2)

	// Lesson 11: Logical Operations
	
	// Lesson 12: If Else Statements
	if s1 > s2 && s1 < 20 {
		fmt.Printf("\n%d is greater than %d and less than 20\n", s1, s2)
	} else if s1 == s2 || s2 < 10 {
		fmt.Printf("\n%d is equal to %d or %d is less than 10\n", s1, s2, s2)
	} else {
		fmt.Printf("\nNeither condition is true\n")
	}
}