package algos

// Binary search  finds if in the sorted slice arr there is the integer target
// time complexity O(log n)
// memory O(N)
// returns (true, index) if value exists
// otherwise (false, -1)
func BinarySearch(target int ,arr []int) (bool, int) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (left +right) / 2
		if arr[middle] < target {
			left = middle + 1
		} else if arr[middle] > target {
			right = middle - 1
		} else{
			return true, middle
		}
	}

	return false, -1
}
