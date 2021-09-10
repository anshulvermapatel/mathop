package calc

func Add(a ...int) int {
	var sum int
	for _, num := range a {
		sum = sum + num
	}
	return sum
}