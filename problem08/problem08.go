package problem08

func IndexOfFirstDuplicate(items []int) int {
	occuranceMap := make(map[int]int)
	for index, value := range items {
		_, EXIST := occuranceMap[value]
		if EXIST {
			return index
		}
		occuranceMap[value]++
	}
	return -1
}
