package demo

// GetFibonacci is to create the Fibonacci list
func GetFibonacci(n int) []int {
	list := []int{1, 1}
	for i := 2; i < n; i++ {
		list = append(list, list[i-2]+list[i-1])
	}
	return list
}
