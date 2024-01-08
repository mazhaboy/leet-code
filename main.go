package main

import (
	"fmt"
	"github.com/leet-code/hash_table"
)

func main() {
	new := hash_table.Constructor()
	new.Add(0)
	new.Add(5)
	new.Add(6)
	new.Add(10)
	fmt.Println(new)
	new.Contains(0)
	new.Remove(0)
	fmt.Println(new.Contains(0))
	fmt.Println(new)

}

func addBinary(a string, b string) string {
	n := Max(len(a), len(b))
	//reverse a, b
	res := ""
	a = reverse(a)
	b = reverse(b)
	i := 0
	memory := 0
	for i < n {
		if len(a) <= i {
			res += string(rune((int(b[i]) - '0' + memory) % 2))
			if int(b[i])+memory >= 2 {
				memory = 1
			} else {
				memory = 0
			}
		} else if len(b) <= i {
			res += string(rune((int(a[i]) + memory) % 2))
			if int(a[i])+memory >= 2 {
				memory = 1
			} else {
				memory = 0
			}
		} else {
			res += string(rune((int(a[i]) + memory + int(b[i])) % 2))
			if int(a[i])+int(b[i])+memory >= 2 {
				memory = 1
			} else {
				memory = 0
			}
		}
		i++

	}
	return res // reverse res
}

func reverse(str string) string {
	res := ""
	for i := range str {
		res = string(str[i]) + res
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
