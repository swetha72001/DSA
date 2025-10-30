package problem

// nums = [1, 2, 3, 4, 5, 6, 7]
// k = 3

func RotateByKposition() {
	nums := []int64{1, 2, 3, 4, 5, 6, 7}
	k := 3
	n := len(nums)
	k = k % n

	reverse(nums, 0, int64(n-1))
	reverse(nums, 0, int64(k-1))
	reverse(nums, int64(k), int64(n-1))
}

func reverse(nums []int64, start, end int64) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
