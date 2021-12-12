package basic

func Sum(values ...int) int {
	result := 0
	for _, value := range values {
		result += value
	}
	return result
}

func Multiply(values ...int) int {
	result := 1
	for _, value := range values {
		result *= value
	}
	return result
}
