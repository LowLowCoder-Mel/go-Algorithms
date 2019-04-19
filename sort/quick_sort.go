package sort

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	quick_sort(arr, 0, len(arr) - 1)
}

func quick_sort(arr []int, start, end int) {
	if start >= end {
		return
	}

	i := partition(arr, start, end)
	quick_sort(arr, start, i - 1)
	quick_sort(arr, i + 1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
