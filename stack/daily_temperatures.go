package stack

import (
	"fmt"
)

type DailyTemperature struct {
	Index       int
	Temperature int
}

type DailyTemperaturesStack struct {
	Stack   map[int]*DailyTemperature
	LastKey int
}

func ConstructorDailyTemperaturesStack() *DailyTemperaturesStack {
	newMyStack := new(DailyTemperaturesStack)
	s := make(map[int]*DailyTemperature, 0)
	newMyStack.Stack = s
	newMyStack.LastKey = 0
	return newMyStack
}

func (this *DailyTemperaturesStack) Push(val *DailyTemperature) {
	this.LastKey++
	this.Stack[this.LastKey] = val
}

func (this *DailyTemperaturesStack) Pop() *DailyTemperature {
	r := &DailyTemperature{0, 0}
	if this.NonEmpty() {
		r = this.Stack[this.LastKey]
		delete(this.Stack, this.LastKey)
		this.LastKey--
	}

	return r
}

func (this *DailyTemperaturesStack) Top() *DailyTemperature {
	if this.NonEmpty() {
		return this.Stack[this.LastKey]
	}

	return &DailyTemperature{
		Index:       0,
		Temperature: 0,
	}
}

func (this *DailyTemperaturesStack) NonEmpty() bool {
	if this.LastKey != 0 {
		return true
	}

	return false
}

func DailyTemperatures(temperatures []int) []int {
	stack := ConstructorDailyTemperaturesStack()
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for temperatures[i] > stack.Top().Temperature && stack.NonEmpty() {
			d := stack.Pop()
			result[d.Index] = i - d.Index
			fmt.Println(result)
		}
		stack.Push(&DailyTemperature{
			Index:       i,
			Temperature: temperatures[i],
		})
	}

	return result
}
