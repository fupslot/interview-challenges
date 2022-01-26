package find

func Duplicates(values []int) []int {
	dups := make(map[int]int)
	var result []int
	for _, v := range values {
		dups[v] += 1
		if dups[v] == 2 {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return []int{-1}
	}

	return result
}

// 0 - n'th
// i % n + size
// this approach works because all elements are in the range
// form 0 to n-1 and values[i] would be greate than n only if
// the value appears twice
func DuplicatesN(values []int) []int {
	size := len(values)
	var result []int

	for i := 0; i < size; i++ {
		idx := values[i] % size
		values[idx] = values[idx] + size
	}

	for i := 0; i < size; i++ {
		if values[i] >= size*2 {
			result = append(result, i)
		}
	}

	return result
}
