package algos

import "testing"

func Test_MergeSort(t *testing.T){
	arr := MergeSort([]int{3,2,4,1,5,8,7,6,9})
	sorted := []int{1,2,3,4,5,6,7,8,9}
	for ix, v := range arr {
		if sorted[ix] != v {
			t.Errorf("Merge sort is borken %v", arr)
		}
	}
}