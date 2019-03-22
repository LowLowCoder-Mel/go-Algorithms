package TwoSum

import "github.com/pkg/errors"

func twoSum(nums []int, target int) ([]int, error){
	hashMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		value, ok := hashMap[key]
		if ok {
			return []int{value, i}, nil
		}

		hashMap[nums[i]] = i
	}

	return []int{0, 0}, errors.New("No two sum solution")
}


