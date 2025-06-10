// operator/operator.go
package operator

// Add trả về tổng của hai số nguyên
func Add(a, b int) int {
	return a + b
}

// Sub trả về hiệu của hai số nguyên
func Sub(a, b int) int {
	return a - b
}

// Mul trả về tích của hai số nguyên
func Mul(a, b int) int {
	return a * b
}

// Div trả về thương (phần nguyên) của hai số nguyên
func Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// Mod trả về phần dư của phép chia hai số nguyên
func Mod(a, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}
