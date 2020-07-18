package main

import "fmt"

func main() {
	x := 1221
	fmt.Println(isPalindrome(x))
}
func isPalindrome(x int) bool {
	l := x
	nbr := 0
	y := 0

	for l != 0 {
		nbr = l % 10
		y = y*10 + nbr
		l = l / 10

	}
	if x < 0 {
		return false
	}
	if x == y {
		return true
	}

	return false

}
