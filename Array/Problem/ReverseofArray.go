package problem

import "log"

func RevreseOfArray() {
	num := []int64{1, 2, 3, 4, 5}
	n := 0
	for j := len(num) - 1; j > len(num)/2; j-- {
		num[j], num[n] = num[n], num[j]
		n++
	}
	log.Println("num>>>>>>>", num)
}
