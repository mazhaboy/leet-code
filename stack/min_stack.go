package stack

type MinStack struct {
	Stack   map[int]int
	LastKey int
}

func Constructor() MinStack {
	newMyStack := new(MinStack)
	s := make(map[int]int, 0)
	newMyStack.Stack = s
	newMyStack.LastKey = 0
	return *newMyStack
}

func (this *MinStack) Push(val int) {
	this.LastKey++
	this.Stack[this.LastKey] = val
}

func (this *MinStack) Pop() {
	if this.NonEmpty() {
		delete(this.Stack, this.LastKey)
		this.LastKey--
	}
}

func (this *MinStack) Top() int {
	if this.NonEmpty() {
		return this.Stack[this.LastKey]
	}

	return 0
}

func (this *MinStack) GetMin() int {
	var min int
	if this.NonEmpty() {
		min = this.Stack[1]
	}
	for i := range this.Stack {
		min = Min(min, this.Stack[i])
	}

	return min
}

func (this *MinStack) NonEmpty() bool {
	if this.LastKey != 0 {
		return true
	}

	return false
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
