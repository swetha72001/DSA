package problem

import "log"

// Find second largest element

func SecondLargest() {
	num := []int64{1, 2, 4, 3, 5}
	//sorting
	n := len(num)
	if n < 2 {
		log.Println("need multiple element")

	}
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	var secondMax int64 = -1
	max := num[n-1]
	for i := n - 2; i >= 0; i-- {
		if num[i] < max {
			secondMax = num[i]
			break
		}
	}
	if secondMax != -1 {
		log.Println("second largest is:", secondMax)

	} else {
		log.Println("no distinct element")

	}
}
