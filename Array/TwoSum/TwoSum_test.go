package TwoSum

import "testing"

func TestTwoSum(t *testing.T) {
	testNum := []int{2, 7, 11, 15}

	result, err := twoSum(testNum, 9)
	if err != nil {
		t.Error(err)
	}
	if result[0] != 0 || result[1] != 1 {
		t.Error("twoNums Error1")
	}

	testNum = []int{2, 7, 11, 15}
	result, err = twoSum(testNum, 26)
	if err != nil {
		t.Error(err)
	}
	if result[0] != 2 || result[1] != 3 {
		t.Error("twoNums Error2")
	}

	testNum = []int{2, 7, 11, 15}
	result, err = twoSum(testNum, 36)
	if err != nil {
		t.Log(err)
	}
}


