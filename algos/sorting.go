package algos

// MergeSort algorithm
// Best    n log n
// Worst   n log n
// Average n log n
// Memory  n
// Stable  Y
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	n := len(arr)/2
	l := MergeSort(arr[:n])
	r := MergeSort(arr[n:])
	return merge(l, r)
}

func merge(left, right []int) []int{
	ret := make([]int, 0, len(left) + len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(ret, right ...)
		}
		if len(right) == 0 {
			return append(ret, left ...)
		}
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}

	return ret
}