package main

import "fmt"

func main() {
	a := 5
	fmt.Println(mySqrt(a))
}

func mySqrt(x int) int {
	i := 0
	for j := x; j > 0; j-- {
		for i = 1; i <= x; i++ {
			if j == i*i {
				return i
			}
		}
	}
	return 0
}
