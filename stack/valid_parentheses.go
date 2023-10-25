package stack

func IsValid(s string) bool {
	stack := NewMyStack()
	closeToOpen := map[uint8]uint8{'}': '{', ')': '(', ']': '['}

	for i := range s {
		_, ok := closeToOpen[s[i]]
		if ok {
			if !stack.Empty() && stack.GetLast() == closeToOpen[s[i]] {
				stack.Pop()
			} else {
				return false
			}
		} else {
			stack.Push(s[i])
		}
	}

	if !stack.Empty() {
		return false
	}

	return true
}
