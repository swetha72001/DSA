package problem

import "log"

func HighFreqencyNumber() {
	str := []int64{4, 4, 1, 2, 1, 1, 2, 4, 4, 4}
	freqMap := make(map[int64]int64)

	for _, v := range str {
		freqMap[v]++
	}
	var max, field int64
	for i, v := range freqMap {

		if v > max {
			max = v
			field = i
		}

	}
	log.Printf("max frequency is %d cocurs %d", field, max)
}
