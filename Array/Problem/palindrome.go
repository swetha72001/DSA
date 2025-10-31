package problem

import "log"

func CheckPalindrome() {
	str := "madam"
	isPalindrome := true
	// var str2 []string
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			isPalindrome = false
		}
	}
	log.Println("palindrome>>>", isPalindrome)
}
