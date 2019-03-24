package MergeSortArray

import "testing"

func TestMergeSortArray(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}

	mergeSortArray(num1, 3, num2, 3)
	t.Log(num1)

	if !compareArray(num1, []int{1, 2, 2, 3, 5, 6}) {
		t.Fail()
	}
}

func compareArray(num1 []int, num2 []int) bool {
	if len(num1) != len(num2) {
		return false
	}
	for i := 0; i < len(num1) - 1; i++ {
		if num1[i] != num2[i] {
			return false
		}
	}

	return true
}
