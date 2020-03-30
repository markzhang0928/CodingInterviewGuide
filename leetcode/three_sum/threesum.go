package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	triplets := make([][]int, 0)
	sort.Ints(nums)

	if len(nums) < 3 {
		return triplets
	}

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		desired := 0 - nums[i]
		start := i + 1
		end := len(nums) - 1
		for start < end {
			if nums[start]+nums[end] == desired {
				triplets = append(triplets, []int{nums[i], nums[start], nums[end]})
				end--
				for start < end && nums[end] == nums[end+1] {
					end--
				}
			} else if nums[start]+nums[end] > desired {
				end--
			} else {
				start++
			}
		}
	}
	return triplets
}
