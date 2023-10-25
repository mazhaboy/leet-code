package stack

type MyStack struct {
	Stack   map[int]uint8
	LastKey int
}

func NewMyStack() MyStack {
	newMyStack := new(MyStack)
	s := make(map[int]uint8, 0)
	newMyStack.Stack = s
	newMyStack.LastKey = 0
	return *newMyStack
}

func (s *MyStack) Push(data uint8) {
	s.LastKey++
	s.Stack[s.LastKey] = data
}

func (s *MyStack) Pop() uint8 {
	data := s.Stack[s.LastKey]
	delete(s.Stack, s.LastKey)
	s.LastKey--
	return data
}

func (s *MyStack) GetLast() uint8 {
	data := s.Stack[s.LastKey]
	return data
}

func (s *MyStack) Empty() bool {
	if s.LastKey == 0 {
		return true
	}

	return false
}
