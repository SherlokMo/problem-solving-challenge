package problem03

import (
	"strconv"
)

func RunLengthEncode(str string) string {
	encode := ""

	for i := 0; i < len(str); {
		counter, character := 1, str[i]
		j := i
		for j < len(str)-1 && str[j] == str[j+1] {
			counter++
			j++
		}

		encode += string(character) + strconv.Itoa(counter)
		i = j + 1
	}

	return encode
}

func RunLengthDecode(str string) string {
	decode := ""

	for i := 0; i < len(str); {
		character := str[i]
		j := i + 1
		// checking if str[j] is numerical between 0-9 using ASCII Code
		// 0 -> 48
		// 9 -> 57
		for j < len(str) && str[j] >= 48 && str[j] <= 57 {
			j++
		}

		occurances, _ := strconv.Atoi(str[i+1 : j])
		for occurances > 0 {
			decode += string(character)
			occurances--
		}

		i = j
	}
	return decode
}
