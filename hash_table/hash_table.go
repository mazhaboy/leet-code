package hash_table

type MyHashSet struct {
	arr []int
}

func Constructor() MyHashSet {
	myHashSet := new(MyHashSet)
	arr := make([]int, 1000001)
	myHashSet.arr = arr
	return *myHashSet
}

func (this *MyHashSet) Add(key int) {
	y := key % 1000001
	this.arr[y] = y
	if key == 0 {
		this.arr[0] = 1
	}
}

func (this *MyHashSet) Remove(key int) {
	y := key % 1000001
	this.arr[y] = 0
}

func (this *MyHashSet) Contains(key int) bool {
	y := key % 1000001

	if key == 0 && this.arr[y] == 1 {
		return true
	} else if this.arr[y] == y && key != 0 {
		return true
	}

	return false
}
