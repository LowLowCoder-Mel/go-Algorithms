package ThreeSum

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	testArray := []int{-1, 0, 1, 2, -1, -4}

	result, err := threeSum(testArray)
	if err != nil {
		t.Fail()
	}

	t.Log(result)
}
