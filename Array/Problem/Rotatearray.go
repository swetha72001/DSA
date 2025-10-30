package problem

import "log"

func RotateArry() {
	nums := []int64{3, 5, 1, 8, 2, 7}
	n := len(nums)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		nums[j], nums[i] = nums[i], nums[j]
	}
	log.Println("reversed array>>>>>>>", nums)
}
