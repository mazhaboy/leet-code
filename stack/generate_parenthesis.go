package stack

import (
	"fmt"
)

type StackStr struct {
	Stack   map[int]string
	LastKey int
}

func ConstructorStack() StackStr {
	newMyStack := new(StackStr)
	s := make(map[int]string, 0)
	newMyStack.Stack = s
	newMyStack.LastKey = 0
	return *newMyStack
}

func (this *StackStr) Push(val string) {
	this.LastKey++
	this.Stack[this.LastKey] = val
}

func (this *StackStr) Pop() string {
	str := ""
	if this.NonEmpty() {
		str = this.Stack[this.LastKey]
		delete(this.Stack, this.LastKey)
		this.LastKey--
	}
	return str
}

func (this *StackStr) Top() string {
	if this.NonEmpty() {
		return this.Stack[this.LastKey]
	}

	return ""
}

func (this *StackStr) NonEmpty() bool {
	if this.LastKey != 0 {
		return true
	}

	return false
}

func (this *StackStr) StirngReverse() string {
	str := ""
	if this.NonEmpty() {
		for i := 1; i <= len(this.Stack); i++ {
			str = str + this.Stack[i]
		}
	}
	return str
}

var (
	s       = ConstructorStack()
	nGlobal int
	res     = make([]string, 0)
)

func GenerateParenthesis(n int) []string {
	backtracking(0, 0, n)
	return res
}

func backtracking(opened, closed, nGlobal int) {
	fmt.Println(s.Stack)
	if opened == nGlobal && closed == nGlobal && opened == closed {
		res = append(res, s.StirngReverse())
		return
	}

	if opened < nGlobal {
		s.Push("(")
		backtracking(opened+1, closed, nGlobal)
		s.Pop()
	}

	if closed < opened {
		s.Push(")")
		backtracking(opened, closed+1, nGlobal)
		s.Pop()
	}
}
