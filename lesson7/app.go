package main

import (
	"bufio"
	"fmt"
	"os"
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

	// Lesson 8: Tìm hiểu về Package fmt - In, Định dạng, Nhập/Xuất dữ liệu
	fmt.Println("\n--- Lesson 8: fmt Package ---")

	// Print: In không xuống dòng
	fmt.Print("Hello ")
	fmt.Print("World")
	fmt.Print("\n")

	// Println: In và xuống dòng
	fmt.Println("Hello World with Println")

	// Printf: In với định dạng
	name := "Go"
	version := 1.21
	fmt.Printf("Language: %s, Version: %.2f\n", name, version)

	// Sprint: Xử lý chuỗi, không in ra màn hình
	message := fmt.Sprint("Hello ", "Go")
	fmt.Println("Sprint result:", message)

	// Sprintln: Xử lý chuỗi với xuống dòng
	messageWithNewline := fmt.Sprintln("Hello", "Go", "with Sprintln")
	fmt.Println("Sprintln result:", messageWithNewline)

	// Sprintf: Xử lý chuỗi với định dạng
	formattedMessage := fmt.Sprintf("Language: %s, Version: %.2f", name, version)
	fmt.Println("Sprintf result:", formattedMessage)

	// Scan: Nhập dữ liệu từ bàn phím (nhập một giá trị)
	var inputName string
	fmt.Print("Nhập tên (1 từ): ")
	fmt.Scan(&inputName)
	
	// Xóa buffer để tránh Scanln bị ăn Enter từ Scan
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	
	// Scanln: Nhập dữ liệu từ bàn phím (nhập cả dòng)
	var inputAge int
	fmt.Print("Nhập tuổi: ")
	fmt.Scanln(&inputAge)
	fmt.Println("Tên bạn nhập:", inputName)
	fmt.Println("Tuổi bạn nhập:", inputAge)

	// Scanf: Nhập dữ liệu với định dạng
	var height float64
	fmt.Print("Nhập chiều cao (m): ")
	fmt.Scanf("%f", &height)
	fmt.Printf("Chiều cao bạn nhập: %.2f m\n", height)
}