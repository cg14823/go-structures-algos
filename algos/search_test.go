package algos

import "testing"

func TestBinarySearch(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t1 := 2
	ans1 := 1

	t2 := 0
	ans2 := -1

	if found, ix := BinarySearch(t1,arr1); !found || ix != ans1 {
		t.Errorf("Should have return (true, %d) instead (%t, %d)", ans1, found, ix)
	}
	if found, ix := BinarySearch(t2,arr1); found || ix != ans2 {
		t.Errorf("Should have return (false, %d) instead (%t, %d)", ans2, found, ix)

	}
}
