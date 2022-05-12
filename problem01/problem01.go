package problem01

func IsPalindrom(str string) bool {
	leftPointer, rightPointer := 0, len(str)-1
	for leftPointer <= rightPointer {
		if str[leftPointer] != str[rightPointer] {
			return false
		}
		leftPointer++
		rightPointer--
	}
	return true
}
