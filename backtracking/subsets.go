package backtracking

type MyArr struct {
	arr [][]int
}

func NewMyArr() *MyArr {
	a := new(MyArr)
	a.arr = make([][]int, 0)
	return a
}

func (a *MyArr) Append(e []int) {
	a.arr = append(a.arr, e)
}

func (a *MyArr) Pop() {
	if len(a.arr) == 0 {
		return
	}
	a.arr = a.arr[:len(a.arr)-1]
}

func Subsets(nums []int) [][]int {
	res := NewMyArr()
	arr := make([]int, 0)
	backtracking(arr, nums, res, 0)
	return res.arr
}

func backtracking(arr, nums []int, res *MyArr, start int) {
	res.Append(arr)
	arr = make([]int, 0)
	for i := start; i < len(nums); i++ {
		arr = append(arr, nums[i])
		backtracking(arr, nums, res, i+1)
		arr = arr[:len(arr)-1]
		backtracking(arr, nums, res, i+1)
	}
}
