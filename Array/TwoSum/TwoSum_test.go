package TwoSum

import "testing"

func TestTwoSum(t *testing.T) {
	testNums := []int{2, 7, 11, 15}
	result := twoSum(testNums, 9)
	if result[0] != 0 || result[1] != 1 {
		t.Error("twoNums Error1")
	}

	testNums = []int{2, 7, 11, 15}
	result = twoSum(testNums, 26)
	if result[0] != 2 || result[1] != 3 {
		t.Error("twoNums Error2")
	}
}


