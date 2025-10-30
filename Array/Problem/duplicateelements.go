package problem

import "log"

//find duplicate element

func DuplicateElements() {
	arr := []int64{1, 2, 1, 3, 4, 5}
	duplicateMap := make(map[int64]bool)
	for _, v := range arr {

		if duplicateMap[v] {
			log.Println("Duplicate element:", v)
		} else {
			duplicateMap[v] = true
		}
	}
}
