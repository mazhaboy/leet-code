package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Deque double ended queue data struct.
type Deque struct {
	Arr []*TreeNode
}

//Constructor init deque.
func Constructor() *Deque {
	newDeque := new(Deque)
	newArr := make([]*TreeNode, 0)
	newDeque.Arr = newArr
	return newDeque
}

//PushFront inserts the element at the beginning.
func (d *Deque) PushFront(e *TreeNode) {
	d.Arr = append([]*TreeNode{e}, d.Arr...)
}

//PushBack adds element at the end.
func (d *Deque) PushBack(e *TreeNode) {
	d.Arr = append(d.Arr, e)
}

//PopFront removes the first element from the deque.
func (d *Deque) PopFront() (*TreeNode, bool) {
	if len(d.Arr) == 0 {
		return &TreeNode{}, false
	}
	firstElement := d.Arr[0]
	d.Arr = d.Arr[1:]
	return firstElement, true
}

//PopBack removes the last element from the deque.
func (d *Deque) PopBack() (*TreeNode, bool) {
	if len(d.Arr) == 0 {
		return &TreeNode{}, false
	}
	lastElement := d.Arr[len(d.Arr)-1]
	d.Arr = d.Arr[:len(d.Arr)-1]
	return lastElement, true
}

//Front gets the front element from the deque.
func (d *Deque) Front() (*TreeNode, bool) {
	if len(d.Arr) == 0 {
		return &TreeNode{}, false
	}
	firstElement := d.Arr[0]
	return firstElement, true
}

//Back gets the last element from the deque.
func (d *Deque) Back() (*TreeNode, bool) {
	if len(d.Arr) == 0 {
		return &TreeNode{}, false
	}
	lastElement := d.Arr[len(d.Arr)-1]
	return lastElement, true
}

//Empty checks whether the deque is empty or not.
func (d *Deque) Empty() bool {
	if len(d.Arr) == 0 {
		return true
	}

	return false
}

//Size determines the number of elements in the deque.
func (d *Deque) Size() int {
	return len(d.Arr)
}
