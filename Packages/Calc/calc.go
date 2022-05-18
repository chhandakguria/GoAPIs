package Calc

func Calculator(a, b int, ch byte) int {

	var c int
	switch ch {
	case '+':
		c = a + b
	case '-':
		c = a - b
	case '*':
		c = a * b
	case '/':
		c = a / b
	}
	return c
}
