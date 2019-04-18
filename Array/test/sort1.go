package main

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	p := partition(arr, start, end)
	quickSort(arr, start, p - 1)
	quickSort(arr, p + 1, end)
}

func partition(arr []int, start, end int) int {
	pvoit := arr[end]
	i := start
	for j := start; j <= end - 1; j++ {
		if arr[j] < pvoit {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func bsearch(arr []int, value int) int {
	low := 0
	high := len(arr)
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] == value {
			return mid
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

//
func bsearch_0(arr []int, value int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid - 1] != value {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func bsearch_1(arr []int, value int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > value {
			high = mid - 1
		} else if arr[mid] < value {
			low = mid + 1
		} else {
			if mid == (len(arr) - 1) || arr[mid + 1] != value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

func bsearch_2(arr []int, value int) int {
	n := len(arr)
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] >= value {
			if mid == 0 || arr[mid - 1] < value {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}

	return -1
}

func bsearch_3(arr []int, value int) int {
	n := len(arr) - 1
	low := 0
	high := n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > value {
			high = mid - 1
		} else {
			if mid == n - 1 || arr[mid + 1] > value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

func main() {

	//test := []int{6, 5, 4, 3, 2, 1, -1}
	//quickSort(test, 0, len(test)-1)
	//log.Println(test)

	//test1 := []int{6, 5, 4, 3, 2, 1}
	//mergeSort(test1, 0, len(test1) - 1)
	//log.Println(test1)

	//test3 := []int{ 1, 2, 3, 4, 5, 6, 7 }
	//result := bsearch(test3,3)
	//print(result)

	test4 := []int{ 1, 2, 2, 3, 3, 4, 5, 6, 7 }
	result := bsearch_3(test4,3)
	print(result)
	
}
