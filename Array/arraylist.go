package main

type Array struct {
	cap int
	count int
	array []int
}

func NewArray(cap int) *Array {
	return &Array{cap: cap, count: 0, array: []int{}}
}

func (arr *Array) find(index int) int {
	if index < 0 || index >= arr.count {
		return -1
	}

	return arr.array[index]
}

func (arr *Array) insert(index int, value int) bool {
	// 检查容量
	if arr.count == arr.cap {
		return false
	}
	// 检查索引
	if index < 0 || index > arr.count {
		return false
	}

	for i := arr.count; i > index; i-- {
		arr.array[i] = arr.array[i - 1]
	}
	arr.array[index] = value
	arr.count++

	return true
}

func (arr *Array) delete(index int) bool {
	if index < 0 || index > arr.count {
		return false
	}

	for i := index+1; i < arr.count; i++ {
		arr.array[i-1] = arr.array[i]
	}
	arr.count--

	return true
}