package main

func majorityElement(nums []int) int {
	count := 0
	vote := nums[0]
	for _, v := range nums {
		// 遇到相同的则票数 + 1，遇到不同的则票数 - 1
		if count == 0 {
			vote = v
			count = 1
		} else if v == vote {
			count++
		} else {
			count--
		}
	}
	return vote
}
