package main

import (
	"fmt"
)

// Fibo
func fibo() func() int {
	a := 0
	b := 1

	return func() int {
		defer func() {
			a, b = b, a+b

		}()
		return a
	}

}

func main() {
	f := fibo()

	fmt.Println(f()) // 1
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 5
	fmt.Println(f()) // 8
	fmt.Println(f()) // 13

}
