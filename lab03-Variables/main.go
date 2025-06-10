package main

import "fmt"

var globalVar = "Tôi là biến toàn cục"

func main() {
	//| Verb  | Ý nghĩa                             | Ví dụ (`Printf`)                 |
	// | ----- | ----------------------------------- | -------------------------------- |
	// | `%T`  | In ra **kiểu dữ liệu**              | `%T` → `string`, `int`, v.v.     |
	// | `%v`  | In **giá trị mặc định**             | `%v` → `Hello`, `42`, `true`     |
	// | `%+v` | In giá trị struct + tên trường      | `%+v` → `{Name: "Hiếu"}`         |
	// | `%#v` | In giá trị dạng Go code             | `%#v` → `main.User{Name:"Hiếu"}` |
	// | `%s`  | In **chuỗi**                        | `%s` → `Hello`                   |
	// | `%d`  | In **số nguyên** (decimal)          | `%d` → `123`                     |
	// | `%f`  | In **số thực**                      | `%f` → `3.140000`                |
	// | `%t`  | In **giá trị boolean**              | `%t` → `true` hoặc `false`       |
	// | `%c`  | In ký tự tương ứng mã ASCII/Unicode | `%c` → `A`                       |
	// | `%b`  | Nhị phân (binary)                   | `%b` → `1101`                    |
	// | `%x`  | Hex (lowercase)                     | `%x` → `2a`                      |
	// | `%p`  | Địa chỉ con trỏ (pointer address)   | `%p` → `0xc000010230`            |

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
