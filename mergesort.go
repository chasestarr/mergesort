package mergesort

// Sort sorts data using mergesort alg
func Sort(data []int) []int {
	return mergesort(data)
}

func mergesort(data []int) []int {
	// base case
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := data[:mid]
	right := data[mid:]

	return merge(mergesort(left), mergesort(right))
}

func merge(a, b []int) []int {
	result := []int{}
	var i int
	var j int
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}

	}

	if i < len(a) {
		result = append(result, a[i:]...)
	}
	if j < len(b) {
		result = append(result, b[j:]...)
	}

	return result
}
