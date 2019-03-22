package ThreeSum

import "errors"

func threeSum(nums []int) ([][]int, error) {
	if len(nums) < 3 {
		return [][]int{}, errors.New("数组个数应该高于3个元素")
	}

	quickSort(nums, 0, len(nums) - 1)

	result := make([][]int, 0)
	for i := 0; i < len(nums) - 1; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		start := i + 1
		end := len(nums) - 1
		for end > start {
			sum := nums[i] + nums[start] + nums[end]
			if sum > 0 {
				end--
			} else if sum < 0 {
				start++
			} else {
				result = append(result, []int{nums[i], nums[start], nums[end]})
				for start < end && nums[start] == nums[start + 1] {
					start++
				}
				for start < end && nums[end] == nums[end - 1] {
					end--
				}
				start++
				end--
			}
		}
	}

	return result, nil
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}

	p := partition(nums, start, end)
	quickSort(nums, start, p - 1)
	quickSort(nums, p + 1, end)
}

func partition(nums []int, start int, end int)  int {
	pivot := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]
	return i
}
