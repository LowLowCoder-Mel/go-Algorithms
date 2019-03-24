package MergeSortArray

func mergeSortArray(num1 []int, m int, num2 []int, n int) {

	i := m - 1
	j := n - 1
	k := m + n - 1

	for ;i >= 0 && j >= 0; k-- {
		if num1[i] > num2[j] {
			num1[k] = num1[i]
			i--
		} else {
			num1[k] = num2[j]
			j--
		}
	}

	for j > 0 {
		num1[k] = num1[i]
		k++
	}
}
