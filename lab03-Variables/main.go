package main

import "fmt"

var globalVar = "Tôi là biến toàn cục"

func main() {
	// 1. var với kiểu rõ ràng
	var a int
	a = 10

	// 2. var có gán giá trị
	var b string = "Hello"

	// 3. var và suy luận kiểu
	var c = true

	// 4. nhiều biến
	var x, y, z = 1, 2, 3

	// 5. khối var
	var (
		name = "Hiếu"
		age  = 21

		married = false
	)

	// 6. :=
	message := "Go is awesome!"

	fmt.Println(a, b, c, x, y, z, name, age, married, message, globalVar)
}
