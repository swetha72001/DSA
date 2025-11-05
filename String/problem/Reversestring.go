package problem

import (
	"log"
	"strings"
)

func ReverseString() {
	str := "Reverse"
	ReverseStr := make([]string, len(str))
	for i, j := 0, len(str)-1; i < len(str); i, j = i+1, j-1 {
		ReverseStr[i] = string(str[j])
	}
	log.Println("reversedstring>>>>>>>>>", strings.Join(ReverseStr, ""))
}
