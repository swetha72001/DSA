package problem

func CountVowelsandConsonants() {
	str := "areuijouf"

	var consonants, vowels int32
	vowelMap := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	for i := 0; i < len(str); i++ {
		if vowelMap[rune(str[i])] {
			vowels++
		} else {
			consonants++
		}
	}

}
