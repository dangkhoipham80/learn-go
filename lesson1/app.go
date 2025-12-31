package main

import (
	"fmt"
)

var city = "TP. Hồ Chí Minh"

var (
	country string
	continent = "Châu Á"
)

func main() {
	var fullName string = "Phạm Đăng Khôi"
	fmt.Println("Full Name:", fullName)

	var isStudent = true
	fmt.Println("Is Student:", isStudent)

	var age int
	age = 21
	fmt.Println("Age:", age)

	var hasChildren int
	fmt.Println("Has Children:", hasChildren)

	phone := "(+84) 123 456 789"
	fmt.Println("Phone:", phone)

	fmt.Println("City:", city)

	var height, width float64
	height = 175.5
	width = 70.0
	fmt.Println("\nHeight:", height)
	fmt.Println("Width:", width)

	math, physics, chemistry := 9, 8, 7
	fmt.Println("\nMath:", math)
	fmt.Println("Physics:", physics)
	fmt.Println("Chemistry:", chemistry)

	country = "Việt Nam"
	fmt.Println("\nCountry:", country)
	fmt.Println("Continent:", continent)
}
