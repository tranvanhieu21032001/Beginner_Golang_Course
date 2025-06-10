package main

import (
	"example/lab04-primitives-operator/operator"
	"fmt"
)

func main() {
	// | Loại          | Kiểu cụ thể                                            | Ví dụ giá trị   |
	// | ------------- | ------------------------------------------------------ | --------------- |
	// | **Số nguyên** | `int`, `int8`, `int16`, `int32`, `int64`               | `42`, `-10`     |
	// | **Số thực**   | `float32`, `float64`                                   | `3.14`, `-2.7`  |
	// | **Số phức**   | `complex64`, `complex128`                              | `2 + 3i`        |
	// | **Boolean**   | `bool`                                                 | `true`, `false` |
	// | **Ký tự**     | `byte` (alias của `uint8`), `rune` (alias của `int32`) | `'A'`, `d'`     |
	// | **Chuỗi**     | `string`                                               | `"Hello"`       |

	var a int = 10
	var b float64 = 3.14
	var c bool = true
	var d string = "GoLang"
	e := 1 == 2

	fmt.Println(a, b, c, d, e)

	// Dùng các toán tử với int
	operator1, operator2 := 29, 9
	fmt.Println("Add:", operator.Add(operator1, operator2))
	fmt.Println("Sub:", operator.Sub(operator1, operator2))
	fmt.Println("Mul:", operator.Mul(operator1, operator2))
	fmt.Println("Div:", operator.Div(operator1, operator2))
	fmt.Println("Mod:", operator.Mod(operator1, operator2))
}
