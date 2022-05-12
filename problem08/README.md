# Problem 08

# Code Sample

```go
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
```

## Complexity Analysis

- Space Complexity is `O(N)` Because I'm using a hashmap that corrosponds to each unique integer in the list, in worst case scenario we are going to have non-repeatable `N` integers on the map.

- Time Complexity is `O(N)` Because we are traversing the array in linear way, in a worst case scenario we are going to scan `N` Integers.

### Other Implementations?

I could've come up with a constant space complexity solution if I we used 2 pointers and bruteforce the items. but this is a non-optimal solution in case of time complexity `O(N^2)`. <br>

sample: 
```go
func IndexOfFirstDuplicate(items []int) int {
	for i := 0; i < len(items) - 1; i++ {
        for j := i + 1; j < len(items); j++ {
            if items[i] == items[j] {
                return j
            }
        }
    }
	return -1
}
```