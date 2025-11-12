package problem

import (
	"log"
	"strings"
)

func SentenceOccurs() {
	str := "this is is a test"
	occurMap := make(map[string]int64)
	words := strings.Fields(str)
	for _, v := range words {
		occurMap[v]++
	}

	for st, v := range occurMap {
		log.Printf(" In sentence there are each %s words %d", st, v)
	}
}
