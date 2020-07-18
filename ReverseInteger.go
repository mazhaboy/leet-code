package main

import "fmt"

func main() {
	x := 2147483647
	fmt.Println(reverse(x))
}
func reverse(x int) int {
	l := x
	nbr := 0
	y := 0

	for l != 0 {
		nbr = l % 10
		y = y*10 + nbr
		l = l / 10

	}
	if y >= -2147483648 && y <= 2147483647 {
		return y
	}
	return 0
}
