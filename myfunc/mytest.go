package myfunc

import "fmt"

func init()  {
	println("init 1111")
}

func init()  {
	println("init 2222")
}

func Hello()  {
	fmt.Printf("hello, %v", "world")
}

func Square(n int) int {
	return n * n
}

func Fibonacci(n int) []int {
	var res = []int{1, 1}
	for i:=2; i < n; i++ {
		res = append(res, res[i-2] + res[i-1])
	}
	return res
}
