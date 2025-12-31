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

	// Lesson 9: Tìm hiểu Go Formatting Verbs
	fmt.Println("\n--- Lesson 9: Go Formatting Verbs ---")

	// Biến để demo
	str := "Go Language"
	num := 42
	floatNum := 3.14159
	isActive := true
	char := 'A'
	ptr := &num
	const typeFormat = "%%T (type): %T\n"

	// %v: Lấy giá trị của biến
	fmt.Printf("%%v (value): %v\n", num)
	fmt.Printf("%%v (value): %v\n", str)

	// %T: Hiển thị kiểu dữ liệu
	fmt.Printf(typeFormat, num)
	fmt.Printf(typeFormat, str)
	fmt.Printf(typeFormat, floatNum)
	fmt.Printf(typeFormat, isActive)

	// %s: String
	fmt.Printf("%%s (string): %s\n", str)

	// %d: Integer (decimal)
	fmt.Printf("%%d (integer): %d\n", num)

	// %.f: Floating-point number
	fmt.Printf("%%.2f (float): %.2f\n", floatNum)
	fmt.Printf("%%.4f (float): %.4f\n", floatNum)

	// %t: Boolean
	fmt.Printf("%%t (boolean): %t\n", isActive)

	// %c: Character (rune)
	fmt.Printf("%%c (character): %c\n", char)

	// %p: Pointer
	fmt.Printf("%%p (pointer): %p\n", ptr)

	// %%: Hiển thị dấu %
	fmt.Printf("%%q (quoted): %q\n", str)
	fmt.Printf("Discount: 50%% off\n")
}