package problem05

func Unique(items []string) []string {
	mmap := make(map[string]int)
	uniqueList := []string{}
	for _, item := range items {
		mmap[item]++
	}
	for k, v := range mmap {
		if v < 2 {
			uniqueList = append(uniqueList, k)
		}
	}

	return uniqueList
}
