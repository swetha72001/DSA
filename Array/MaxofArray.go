package main

import "log"

// Problem Statement:
// Given an array of integers, find and return the maximum element in the array.

// Example Input:
// nums = [3, 5, 1, 8, 2, 7]

func MaxOfArray() int64 {
	nums := []int64{3, 5, 1, 8, 2, 7}
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	log.Println("max>>>", max)
	return max
}

func main() {
	MaxOfArray()
}
