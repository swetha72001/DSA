package problem

import "log"

func Frequencyofcharacters() {
	str := "aleseka"

	freqMap := make(map[rune]int64)

	for _, v := range str {
		freqMap[v]++
	}

	for i, v := range freqMap {
		log.Printf("freq of characyers is %c,%d", i, v)
	}
}
