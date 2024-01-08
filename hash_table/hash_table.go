package hash_table

type MyHashMap struct {
	arr [1000001]int
}

func ConstructorMyHashMap() MyHashMap {
	myHashMap := new(MyHashMap)
	arr := [1000001]int{}
	for i := range arr {
		arr[i] = -1
	}
	myHashMap.arr = arr
	return *myHashMap
}

func (this *MyHashMap) Put(key int, value int) {
	y := key % 1000001
	this.arr[y] = value
}

func (this *MyHashMap) Get(key int) int {
	y := key % 1000001
	return this.arr[y]
}

func (this *MyHashMap) Remove(key int) {
	y := key % 1000001
	this.arr[y] = -1
}
