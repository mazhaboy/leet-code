package tree

import (
	"fmt"
)

type MyArr struct {
	arr []int
}

func kthSmallest(root *TreeNode, k int) int {
	arr := &MyArr{arr: make([]int, 0)}
	orderedArray(root, arr)
	fmt.Println(arr)
	if len(arr.arr) == 0 {
		return 0
	}
	return arr.arr[k-1]
}

func orderedArray(node *TreeNode, arr *MyArr) {
	if node == nil {
		return
	}
	if node.Left != nil {
		orderedArray(node.Left, arr)
	}
	arr.arr = append(arr.arr, node.Val)
	if node.Right != nil {
		orderedArray(node.Right, arr)
	}
}
